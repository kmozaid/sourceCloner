package main

import (
	handlers "github.com/mozaidk/sourceCloner/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/authorize", handlers.AuthorizeHandler)
	http.HandleFunc("/welcome", handlers.WelcomeHandler)
	http.HandleFunc("/clone", handlers.CloneHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
