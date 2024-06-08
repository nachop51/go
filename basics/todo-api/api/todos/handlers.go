package todos

import (
	"api/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var TodoList []models.Todo = []models.Todo{
	{
		ID:        1,
		Title:     "First Todo",
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func TodoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "Invalid ID"})
			return
		}

		var todo *models.Todo

		for _, t := range TodoList {
			if t.ID == idInt {
				todo = &t
				break
			}
		}

		if todo == nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": "Todo not found"})
			return
		}

		ctx := context.WithValue(r.Context(), "todo", todo)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, TodoList)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {

	// todoData := render.Bind(r, &models.Todo{})

	// todo := models.Todo{
	// 	ID:        len(TodoList) + 1,
	// 	Title:     "New Todo",
	// 	Completed: false,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// TodoList = append(TodoList, todo)

	// render.Status(r, http.StatusCreated)

	// render.JSON(w, r, todo)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value("todo").(*models.Todo)

	render.JSON(w, r, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value("todo").(*models.Todo)

	for i, t := range TodoList {
		if t.ID == todo.ID {
			TodoList = append(TodoList[:i], TodoList[i+1:]...)
			break
		}
	}

	render.NoContent(w, r)
}
