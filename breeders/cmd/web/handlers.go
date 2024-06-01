package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {

	// * How to use the nil field in the app.render example
	// data := &templateData{
	// 	Data: map[string]interface{}{
	// 		"pageTitle": "Home",
	// 	},
	// }
	// app.render(w, "home.page.gohtml", data)

	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	app.render(w, page+".page.gohtml", nil)
}
