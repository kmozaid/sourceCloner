package service

import (
	"github.com/mozaidk/sourceCloner/model"
	"github.com/mozaidk/sourceCloner/service/github"
)

var gitHubCloneService = github.CloneService{}

func GetRepositories(accessToken string) model.RepositoryList {
	return gitHubCloneService.GetRepositories(accessToken)
}

func CloneRepository(url string, name string, dir string, token string) model.CloneResponse {
	return gitHubCloneService.CloneRepository(url, name, dir, token)
}
