package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:easy@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err)})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("You connected to your database.")

	defer db.Close()

	router := mux.NewRouter()
	router.Use(jsonMiddleware)

	// Insert
	// insert, err := db.Query("INSERT INTO users (id, name) VALUES (2, 'John')")
	//
	// update, err := db.Query("UPDATE users SET name = 'Jloooohn' WHERE id = 2")
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer update.Close()
	//
	// defer insert.Close()
	// selects, err := db.Query("SELECT * FROM users")
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer selects.Close()
	//
	// var results []User
	//
	// for selects.Next() {
	// 	var user User
	//
	// 	err = selects.Scan(&user.id, &user.name)
	//
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	results = append(results, user)
	// }
	//
	// fmt.Println(results)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/insert", insertHandler).Methods("post")
	router.HandleFunc("/read", readHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func insertHandler(w http.ResponseWriter, r *http.Request) {

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insert, err := db.Query("INSERT INTO users (id, name) VALUES ($1, $2)", newUser.ID, newUser.Name)

	if err != nil {
		panic(err)
	}

	defer insert.Close()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User inserted successfully"))
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var results []User

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Name)

		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	jsonData, err := json.Marshal(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
