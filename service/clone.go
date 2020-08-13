package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mozaidk/sourceCloner/model"
	"net/http"
	"os"
)

type CloneService interface {
	getRepositories(accessToken string) model.RepositoryList
	cloneRepository(url string, name string, dir string, token string) model.CloneResponse
}

type GitHubCloneService struct {
}

func (g GitHubCloneService) getRepositories(accessToken string) model.RepositoryList {
	reposReqURL := fmt.Sprintf("https://api.github.com/user/repos?access_token=%s", accessToken)
	res, err := http.Get(reposReqURL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
	}
	defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	var repos []model.Repository
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
	}
	return model.RepositoryList{AccessToken: accessToken, Repositories: repos}
}

func (g GitHubCloneService) cloneRepository(url string, name string, dir string, token string) model.CloneResponse {
	fmt.Printf("git clone %s %s \n", url, dir)
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		Auth: &githttp.BasicAuth{
			Username: "something", // yes, this can be anything except an empty string
			Password: token,
		},
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Fprintf(os.Stdout, "Could not clone repository: %v", err)
		return model.CloneResponse{Status: "Failed: " + err.Error()}
	}
	if repo != nil {
		fmt.Printf("Cloned %s", name)
	}

	return model.CloneResponse{Status: "Succeed"}
}

var cloneService CloneService = GitHubCloneService{}

func GetRepositories(accessToken string) model.RepositoryList {
	return cloneService.getRepositories(accessToken)
}

func CloneRepository(url string, name string, dir string, token string) model.CloneResponse {
	return cloneService.cloneRepository(url, name, dir, token)
}
