package main

import (
	"breeders/pets"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {

	// * How to use the nil field in the app.render example
	// data := &templateData{
	// 	Data: map[string]interface{}{
	// 		"pageTitle": "Home",
	// 	},
	// }
	// app.render(w, "home.page.html", data)

	app.render(w, "home.page.html", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	if page == "favicon.ico" {
		http.ServeFile(w, r, "./static/favicon.ico")
		return
	}

	app.render(w, page+".page.html", nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.html", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	dog, err := pets.NewPetFromAbstractFactory("dog")

	if err != nil {
		t.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	cat, err := pets.NewPetFromAbstractFactory("cat")

	if err != nil {
		t.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.Models.DogBreed.All()

	if err != nil {
		t.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	t.WriteJSON(w, http.StatusOK, dogBreeds)
}
