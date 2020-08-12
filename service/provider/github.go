package provider

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mozaidk/sourceCloner/config"
	"github.com/mozaidk/sourceCloner/model"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"os"
	"strings"
)

var conf = &oauth2.Config{
	ClientID:     config.ServiceConf.ClientId,
	ClientSecret: config.ServiceConf.ClientSecret,
	Endpoint:     github.Endpoint,
	RedirectURL:  config.ServiceConf.RedirectUri,
	Scopes:       config.ServiceConf.Scopes,
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

func GetRepositories(accessToken string) model.RepositoryList {
	reqURL := fmt.Sprintf("https://%s@api.github.com/user", accessToken)
	res, err := http.Get(reqURL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
	}
	defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	var user model.User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
	}

	reposReqURL := strings.Replace(user.ReposUrl, "api.", accessToken+"@api.", 1)
	res, err = http.Get(reposReqURL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
	}
	defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	var repos []model.Repository
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
	}
	return model.RepositoryList{Repositories: repos}
}

func CloneRepository(url string, name string, dir string, token string) model.CloneResponse {
	fmt.Printf("git clone %s %s \n", url, dir)

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		Auth: &githttp.BasicAuth{
			Username: "", // yes, this can be anything except an empty string
			Password: token,
		},
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Fprintf(os.Stdout, "Could not clone repository: %v", err)
		return model.CloneResponse{Status: "Failed"}
	}
	if repo != nil {
		fmt.Printf("Cloned %s", name)
	}

	return model.CloneResponse{Status: "Succeed"}
}
