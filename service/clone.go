package service

import (
	"github.com/mozaidk/sourceCloner/model"
	"github.com/mozaidk/sourceCloner/service/provider"
)

func AuthorizeURL() string {
	return provider.AuthorizeURL()
}

func AccessToken(code string) string {
	return provider.AccessToken(code)
}

func GetRepositories(accessToken string) model.RepositoryList {
	return provider.GetRepositories(accessToken)
}

func CloneRepository(url string, name string, dir string, token string) model.CloneResponse {
	return provider.CloneRepository(url, name, dir, token)
}
