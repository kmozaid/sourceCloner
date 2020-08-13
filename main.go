package main

import (
	"fmt"
	"github.com/mozaidk/sourceCloner/config"
	handlers "github.com/mozaidk/sourceCloner/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/authorize", handlers.AuthorizeHandler)
	http.HandleFunc("/welcome", handlers.WelcomeHandler)
	http.HandleFunc("/clone", handlers.CloneHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ServiceConf.Port), nil))
}
