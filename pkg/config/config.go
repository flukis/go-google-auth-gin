package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthCfg = &oauth2.Config{
	Endpoint: google.Endpoint,
	Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
}

const GoogleOauthUrlApi = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func InitGoogleConfig() {
	GoogleOauthCfg.ClientID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	GoogleOauthCfg.ClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	GoogleOauthCfg.RedirectURL = os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")
}
