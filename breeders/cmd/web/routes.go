package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	var mux = chi.NewRouter()

	// Middlewares for the router
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Routes definition starts here
	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	return mux
}
