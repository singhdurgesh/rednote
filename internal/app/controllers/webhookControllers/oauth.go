package webhookControllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/singhdurgesh/rednote/internal/app/services"
	"github.com/singhdurgesh/rednote/internal/constants"
	"github.com/singhdurgesh/rednote/pkg/oauth/googleOAuth"
)

type OAuthController struct{}

var userService = new(services.UserService)
var authService = new(services.AuthService)

func (webhookController *OAuthController) GoogleLogin(ctx *gin.Context) {
	code := ctx.Query("code")

	// 1. Verify the Code and Get User Information

	data, err := googleOAuth.VerifyOAuthCode(ctx, code, ctx.Query("error"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	// 2. From the Details, fetch if the user exists with Verified Email or Create One
	email := data["email"].(string)
	email_verified := data["verified_email"].(bool)

	user := &models.User{}
	res := app.Db.Find(user, "email = ? AND email_verified = ?", email, email_verified)

	if res.Error != nil || res.RowsAffected == 0 {
		name := data["name"].(string)
		user = userService.CreateUser(map[string]interface{}{"name": name, "email": email, "email_verified": email_verified})
		if user == nil {
			ctx.JSON(http.StatusUnauthorized, "User Couldn't be created")
		}
	}

	// 3. Generate JWT Token for that User with flag as GoogleOauth Logged in
	authToken := authService.GenerateJwtToken(user, constants.GOOGLE_AUTH_MODE)

	ctx.JSON(http.StatusOK, gin.H{"token": authToken, "user": user})
}
