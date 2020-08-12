package handler

import (
	"encoding/json"
	"github.com/mozaidk/sourceCloner/config"
	"github.com/mozaidk/sourceCloner/service"
	"net/http"
)

func CloneHandler(w http.ResponseWriter, r *http.Request) {
	ParseForm(w, r)

	repoURL := r.FormValue("repoURL")
	repoName := r.FormValue("repoName")
	repoPath := config.ServiceConf.CloneDir + "/" + repoName
	accessToken := r.FormValue("access_token")

	result := service.CloneRepository(repoURL, repoName, repoPath, accessToken)

	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(res)
}
