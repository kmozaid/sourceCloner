package service

import (
	"github.com/mozaidk/sourceCloner/config"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type OAuthService interface {
	authorizeURL() string
	accessToken(code string) string
}

type GitHubOAuthService struct {
	conf oauth2.Config
}

func (g GitHubOAuthService) authorizeURL() string {
	return g.conf.AuthCodeURL("state")
}

func (g GitHubOAuthService) accessToken(code string) string {
	token, err := g.conf.Exchange(context.Background(), code)
	if err != nil {
		return ""
	}
	return token.AccessToken
}

var gitHubOAuthService OAuthService = GitHubOAuthService{
	conf: oauth2.Config{
		ClientID:     config.ServiceConf.ClientId,
		ClientSecret: config.ServiceConf.ClientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  config.ServiceConf.RedirectUri,
		Scopes:       config.ServiceConf.Scopes,
	},
}

func AuthorizeURL() string {
	return gitHubOAuthService.authorizeURL()
}

func AccessToken(code string) string {
	return gitHubOAuthService.accessToken(code)
}
