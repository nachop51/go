package todos

import (
	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", ListTodos)
	router.Post("/", CreateTodo)

	router.Route("/{id}", func(r chi.Router) {
		r.Use(TodoCtx)
		r.Get("/", GetTodo)
		r.Put("/", UpdateTodo)
		r.Delete("/", DeleteTodo)
	})

	return router
}
