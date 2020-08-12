package handlers

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"net/http"
	"os"
)

type CloneResponse struct {
	Response string
}
func CloneHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	accessToken := r.FormValue("access_token")
	repoURL := r.FormValue("repoURL")
	repoName := r.FormValue("repoName")

	repoPath := "/tmp/" + repoName

	fmt.Printf("git clone %s %s \n", repoURL, repoPath)

	repo, error := git.PlainClone(repoPath, false, &git.CloneOptions{
		Auth: &githttp.BasicAuth{
			Username: "", // yes, this can be anything except an empty string
			Password: accessToken,
		},
		URL:      repoURL,
		Progress: os.Stdout,
	})

	if error != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if repo != nil {
		fmt.Printf("Cloned %s", repoName)
	}
	w.Write([]byte("{\"response\": \"successful\"}"))
}