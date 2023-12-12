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

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create pages cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}