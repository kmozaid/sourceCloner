package main

import (
	"encoding/json"
	"fmt"
	git "github.com/go-git/go-git/v5"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
var validPath = regexp.MustCompile("^/(edit|save|view)$*")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
*/

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

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type CloneResponse struct {
	Response string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/login/oauth/authorize?client_id=21562168ca87f2e900ef&scope=read:userr%20repo", http.StatusSeeOther)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")

	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		"21562168ca87f2e900ef", "a07b152b655155f9878f212114d6901cf20c5f5a", code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	// We set this header since we want the response as JSON
	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	var t OAuthAccessResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// Finally, send a response to redirect the user to the "welcome" page with the access token
	w.Header().Set("Location", "/welcome?access_token="+t.AccessToken)
	w.WriteHeader(http.StatusFound)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
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

func cloneHandler(w http.ResponseWriter, r *http.Request) {
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

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/welcome.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/authorize", authorizeHandler)
	http.HandleFunc("/oauth/redirect", callbackHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/clone", cloneHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
