package service

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var conf = &oauth2.Config{
	ClientID:     "21562168ca87f2e900ef",
	ClientSecret: "a07b152b655155f9878f212114d6901cf20c5f5a",
	Endpoint:     github.Endpoint,
	RedirectURL:  "http://localhost:8080/oauth/redirect",
	Scopes:       []string{"read:user", "repo"},
}

func AuthorizeURL() string {

	return conf.AuthCodeURL("state")
}

func AccessToken(code string) string {
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return ""
	}
	return token.AccessToken
}
