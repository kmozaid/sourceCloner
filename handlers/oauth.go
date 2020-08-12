package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/login/oauth/authorize?client_id=21562168ca87f2e900ef&scope=read:userr%20repo", http.StatusSeeOther)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
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
