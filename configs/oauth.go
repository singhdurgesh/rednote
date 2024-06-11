package configs

type GoogleOAuth struct {
	ClientId       string
	ClientSecret   string
	Scopes         []string
	AuthUrl        string
	AccessTokenUrl string
	RedirectUrl    string
}

type OAuth struct {
	GoogleOAuth GoogleOAuth
}
