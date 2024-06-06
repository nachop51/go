package main

import (
	"breeders/pets"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	render.JSON(w, r, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.html", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	dog, err := pets.NewPetFromAbstractFactory("dog")

	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	cat, err := pets.NewPetFromAbstractFactory("cat")

	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	dogBreeds, err := app.App.Models.DogBreed.All()

	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("Golden Retriever").
		SetMinWeight(55).
		SetMaxWeight(75).
		SetWeight(65).
		SetDescription("Friendly and intelligent").
		SetLifeSpan("10-12 years").
		SetGeographicOrigin("Scotland").
		SetColor("Golden").
		SetAge(2).
		SetAgeEstimated(false).
		Build()

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
	}

	render.JSON(w, r, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("Siamese").
		SetMinWeight(6).
		SetMaxWeight(14).
		SetWeight(10).
		SetDescription("Affectionate and social").
		SetLifeSpan("8-12 years").
		SetGeographicOrigin("Thailand").
		SetColor("Cream, fawn, cinnamon, chocolate, blue").
		SetAge(3).
		SetAgeEstimated(false).
		Build()

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
	}

	render.JSON(w, r, p)
}
