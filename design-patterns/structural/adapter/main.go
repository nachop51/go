package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type ToDo struct {
	UserId    int    `json:"userId" xml:"userId"`
	Id        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type DataInterface interface {
	GetData() (*ToDo, error)
}

type RemoteService struct {
	Remote DataInterface
}

func (r *RemoteService) CallRemoteService() (*ToDo, error) {
	return r.Remote.GetData()
}

type JSONBackend struct{}

func (j *JSONBackend) GetData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var todo ToDo
	err = json.NewDecoder(resp.Body).Decode(&todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

type XMLBackend struct{}

func (x *XMLBackend) GetData() (*ToDo, error) {
	// XML implementation
	xmlFile := `<?xml version="1.0" encoding="UTF-8"?>
	<root>
		<userId>2</userId>
		<id>2</id>
		<title>lorem ipsum dolor</title>
		<completed>true</completed>
	</root>`

	var todo ToDo

	err := xml.Unmarshal([]byte(xmlFile), &todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func main() {
	// Without adapter
	todo := getRemoteData()
	fmt.Printf("todo: %#v\n", todo)

	// With adapter JSON
	service := RemoteService{
		Remote: &JSONBackend{},
	}

	todo, err := service.CallRemoteService()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("todo: %#v\n", todo)

	// With adapter XML
	service = RemoteService{
		Remote: &XMLBackend{},
	}

	todo, err = service.CallRemoteService()

	if err != nil {
		log.Fatal(err)
	}

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
