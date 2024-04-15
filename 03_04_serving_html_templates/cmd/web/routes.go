package main

import (
	"github.com/bmizerany/pat"
	"go_go/03_04_serving_html_templates/pkg/config"
	"go_go/03_04_serving_html_templates/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
