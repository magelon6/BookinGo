package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/magelon6/bookinGo/pkg/config"
	"github.com/magelon6/bookinGo/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

  mux := chi.NewRouter()

  mux.Use(middleware.Recoverer)
  mux.Use(NoSurf)
  mux.Use(SessionLoad)

  mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
  mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

  mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

  return mux
}