package service

import "github.com/mozaidk/sourceCloner/service/github"

var gitHubOAuthService = github.AuthService{}

func AuthorizeURL() string {
	return gitHubOAuthService.AuthorizeURL()
}

func AccessToken(code string) string {
	return gitHubOAuthService.AccessToken(code)
}
