package main

import (
	handlers "github.com/mozaidk/sourceCloner/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/authorize", handlers.AuthorizeHandler)
	http.HandleFunc("/oauth/redirect", handlers.CallbackHandler)
	http.HandleFunc("/welcome", handlers.WelcomeHandler)
	http.HandleFunc("/clone", handlers.CloneHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
