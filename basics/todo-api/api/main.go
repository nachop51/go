package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
)

func main() {
	port := flag.String("port", "3333", "The address to listen on for HTTP requests.")
	genDocRoutes := flag.Bool("genDocRoutes", false, "Generate API documentation routes")
	printRoutes := flag.Bool("printRoutes", false, "Print all available routes")

	flag.Parse()

	router := CreateNewRouter()

	if *printRoutes {

		walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			log.Printf("%s %s\n", method, route)
			return nil
		}

		if err := chi.Walk(router, walkFunc); err != nil {
			log.Panicf("Logging err: %s\n", err.Error())
		}
	}

	if *genDocRoutes {
		// docgen.PrintRoutes(router)
		fmt.Println(docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
			ProjectPath: "github.com/Nachop51/go/basics/todoapi",
			Intro:       "Welcome to the API documentation",
			URLMap:      map[string]string{"GET": "https://example.com"},
		}))
	}

	fmt.Println("Starting server on", ":"+*port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
