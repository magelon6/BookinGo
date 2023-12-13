package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/magelon6/bookinGo/pkg/config"
	"github.com/magelon6/bookinGo/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
  // mux := pat.New()

  // mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
  // mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

  mux := chi.NewRouter()

  mux.Use(middleware.Recoverer)
  mux.Use(NoSurf)

  mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
  mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

  return mux
}