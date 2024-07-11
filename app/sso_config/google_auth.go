package sso_config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/api/v1/auth/sso/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
