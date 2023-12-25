package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:easy@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database.")

	defer db.Close()

	// Insert
	// insert, err := db.Query("INSERT INTO users (id, name) VALUES (2, 'John')")

	// update, err := db.Query("UPDATE users SET name = 'Jloooohn' WHERE id = 2")

	// if err != nil {
	// 	panic(err)
	// }

	// defer update.Close()

	// defer insert.Close()
	selects, err := db.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	type User struct {
		id   int
		name string
	}

	var results []User

	for selects.Next() {
		var user User

		err = selects.Scan(&user.id, &user.name)

		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	fmt.Println(results)
}
