package main

import (
	"api/api/todos"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func CreateNewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60),
		middleware.RedirectSlashes,
	)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"status": "ok"})
	})

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/todos", todos.Routes())
	})

	return router
}
