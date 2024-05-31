package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port string = ":4000"

type appConfig struct {
	useCache bool
}

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}
	//
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//
	// })

	flag.BoolVar(&app.config.useCache, "cache", false, "Enable template cache")
	flag.Parse()

	fmt.Printf("Template cache is enabled: %v\n", app.config.useCache)

	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Printf("Starting server on port http://localhost%s\n", port)

	// err := http.ListenAndServe(port, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
