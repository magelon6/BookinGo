package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/magelon6/bookinGo/pkg/config"
	"github.com/magelon6/bookinGo/pkg/handlers"
	"github.com/magelon6/bookinGo/pkg/render"
)

const portNumber = ":8080"

// Main app func
func main() {

	var app config.AppConfig

	tc, err := render.CreatePagesCache()
	if err != nil {
		log.Fatal("cannot create pages cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}