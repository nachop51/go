package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ToDo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	todo := getRemoteData()

	fmt.Printf("todo: %#v\n", todo)
}

func getRemoteData() *ToDo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var todo ToDo
	err = json.NewDecoder(resp.Body).Decode(&todo)

	if err != nil {
		log.Fatalln(err)
	}

	return &todo
}
