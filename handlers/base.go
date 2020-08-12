package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type User struct {
	UserName string `json:"login"`
	ReposUrl string `json:"repos_url"`
}

type Repository struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	CloneUrl string `json:"clone_url"`
}

type Data struct {
	Repositories []Repository
}

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/welcome.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	accessToken := r.FormValue("access_token")
	reqURL := fmt.Sprintf("https://%s@api.github.com/user", accessToken)

	res, err := http.Get(reqURL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	var user User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	reposReqURL := strings.Replace(user.ReposUrl, "api.", accessToken+"@api.", 1)
	res, err = http.Get(reposReqURL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	var repos []Repository
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	data := &Data{repos}
	renderTemplate(w, "welcome", data)
}