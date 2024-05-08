package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type OAuth struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}

func newOAuth() *OAuth {
	return &OAuth{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectUrl:  os.Getenv("OAUTH_GOOGLE_REDIRECT_URL"),
	}
}

func (c OAuth) GenerateOAuthTokenRequestURL(code string) string {
	baseUrl := "https://oauth2.googleapis.com/token"
	clientId := c.ClientID
	clientSecret := c.ClientSecret
	redirectUrl := c.RedirectUrl
	scopeOpenId := "openid"
	scopeProfile := "https://www.googleapis.com/auth/userinfo.profile"
	scopeEmail := "https://www.googleapis.com/auth/userinfo.email"
	scope := scopeOpenId + " " + scopeProfile + " " + scopeEmail
	requestUrl := "%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s&scope=%s&grant_type=authorization_code"

	tokenURL := fmt.Sprintf(
		requestUrl,
		baseUrl,
		clientId,
		clientSecret,
		url.QueryEscape(code),
		url.QueryEscape(redirectUrl),
		strings.ReplaceAll(url.QueryEscape(scope), "+", "%20"),
	)

	return tokenURL
}
