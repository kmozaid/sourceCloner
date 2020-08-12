package handler

import (
	"github.com/mozaidk/sourceCloner/service"
	"net/http"
)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, service.AuthorizeURL(), http.StatusSeeOther)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	ParseForm(w, r)

	// Get the access token.
	accessToken := service.AccessToken(r.FormValue("code"))
	//fmt.Printf("Access Token from Service %s\n", accessToken)

	// TODO: store token in session or something else.
	w.Header().Set("Location", "/welcome?access_token="+accessToken)
	w.WriteHeader(http.StatusFound)
}
