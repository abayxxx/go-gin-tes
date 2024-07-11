package sso_config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/api/v1/auth/sso/callback",
	ClientID:     "100327984788-oao5ke5cn1rvao2dsggbiidt0ompljkn.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
