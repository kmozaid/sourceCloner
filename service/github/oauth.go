package github

import (
	"github.com/mozaidk/sourceCloner/config"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type AuthService struct {
}

var conf = &oauth2.Config{
	ClientID:     config.ServiceConf.ClientId,
	ClientSecret: config.ServiceConf.ClientSecret,
	Endpoint:     github.Endpoint,
	RedirectURL:  config.ServiceConf.RedirectUri,
	Scopes:       config.ServiceConf.Scopes,
}

func (g AuthService) AuthorizeURL() string {
	return conf.AuthCodeURL("state")
}

func (g AuthService) AccessToken(code string) string {
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return ""
	}
	return token.AccessToken
}
