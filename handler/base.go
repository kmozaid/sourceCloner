package handler

import (
	"fmt"
	"github.com/mozaidk/sourceCloner/service"
	"html/template"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, service.AuthorizeURL(), http.StatusSeeOther)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	ParseForm(w, r)

	accessToken := service.AccessToken(r.FormValue("code"))

	repositoryList := service.GetRepositories(accessToken)

	renderTemplate(w, "welcome", repositoryList)
}

var templates = template.Must(template.ParseFiles("template/index.html", "template/welcome.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ParseForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
