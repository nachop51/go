package main

import (
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any // or interface{}
}

func (app *application) render(w http.ResponseWriter, name string, td *templateData) {
	var tmpl *template.Template

	if app.config.useCache {
		if value, ok := app.templateMap[name]; ok {
			tmpl = value
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(name)

		if err != nil {
			log.Printf("Error parsing template: %s", err)
		}

		log.Printf("Adding template to cache: %s", name)

		tmpl = newTemplate
	}

	if err := tmpl.ExecuteTemplate(w, name, td); err != nil {
		log.Printf("Error executing template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if td == nil {
		td = &templateData{}
	}
}

func (app *application) buildTemplateFromDisk(name string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/head.partial.gohtml",
		"./templates/partials/navbar.partial.gohtml",
		"./templates/partials/footer.partial.gohtml",
		"./templates/" + name,
	}

	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		return nil, err
	}

	app.templateMap[name] = tmpl

	return tmpl, nil
}
