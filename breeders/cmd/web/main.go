package main

import (
	"breeders/adapters"
	"breeders/config"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const port string = ":4000"

type appConfig struct {
	useCache bool
	dsn      string
}

type application struct {
	// Add the template cache to the application struct
	templateMap map[string]*template.Template
	// Another struct that contains a bunch of configuration settings for the app
	config appConfig
	// Add the application struct to the application struct
	App *config.Application
	// Add the RemoteService struct to the application struct
	catService *adapters.RemoteService
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
	flag.StringVar(&app.config.dsn, "dsn", "postgres://postgres:postgres@localhost:5432/breeders?sslmode=disable", "Postgres data source name")
	flag.Parse()

	db, err := initPostgres(app.config.dsn)

	if err != nil {
		log.Panic(err)
	}

	app.App = config.New(db)

	// service := &adapters.RemoteService{
	// 	Remote: &adapters.JSONBackend{},
	// }

	service := &adapters.RemoteService{
		Remote: &adapters.XMLBackend{},
	}

	app.catService = service

	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Printf("Template cache is enabled: %v\n", app.config.useCache)

	fmt.Printf("Starting server on port http://localhost%s\n", port)

	// err := http.ListenAndServe(port, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
