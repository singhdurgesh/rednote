package googleOAuth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/singhdurgesh/rednote/cmd/app"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetSSOUrl() (string, error) {
	authConfig := app.Config.OAuth.GoogleOAuth

	if authConfig.AuthUrl == "" || authConfig.ClientId == "" || authConfig.RedirectUrl == "" {
		return "", errors.New("invalid OAuth configuration")
	}

	URL, err := url.Parse(authConfig.AuthUrl)
	if err != nil {
		return "", err
	}

	parameters := url.Values{}
	parameters.Add("client_id", authConfig.ClientId)
	parameters.Add("scope", strings.Join(authConfig.Scopes, " "))
	parameters.Add("redirect_uri", authConfig.RedirectUrl)
	parameters.Add("response_type", "code")
	parameters.Add("state", generateState()) // generateState should create a random string
	URL.RawQuery = parameters.Encode()

	return URL.String(), nil
}

func VerifyOAuthCode(ctx context.Context, code string, error_reason string) (map[string]interface{}, error) {
	if code == "" {
		message := "code not found."
		app.Logger.Info("Code Not Found.")

		if error_reason == "access_denied" {
			message = message + ", user has denied permission."
		}

		app.Logger.Info(error_reason)
		return nil, errors.New(message)
	}

	oauthConfig := &oauth2.Config{
		ClientID:     app.Config.OAuth.GoogleOAuth.ClientId,
		ClientSecret: app.Config.OAuth.GoogleOAuth.ClientSecret,
		RedirectURL:  app.Config.OAuth.GoogleOAuth.RedirectUrl,
		Scopes:       app.Config.OAuth.GoogleOAuth.Scopes,
		Endpoint:     google.Endpoint,
	}

	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		app.Logger.Error("Token exchange failed: ", err)
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		app.Logger.Error("Failed to get user info: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		app.Logger.Error("Non-OK HTTP status: ", resp.StatusCode)
		return nil, errors.New("failed to get user info")
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Logger.Error("Failed to read response body: ", err)
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		app.Logger.Error("Failed to unmarshal JSON: ", err)
		return nil, errors.New("invalid JSON formatted data")
	}

	return data, nil
}

// generateState generates a random string to be used as the state parameter for CSRF protection
func generateState() string {
	// Implement a secure random string generator
	return "secureRandomString"
}
