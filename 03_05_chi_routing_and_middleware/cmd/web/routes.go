package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_go/03_05_chi_routing_and_middleware/pkg/config"
	"go_go/03_05_chi_routing_and_middleware/pkg/handlers"

	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
