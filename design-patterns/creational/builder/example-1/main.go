package main

import "fmt"

// Builder and Fluent Interface
// The builder pattern is a design pattern that allows for the step-by-step creation of complex objects using the correct sequence of actions.
// The fluent interface is an object-oriented API that provides a more readable and concise code.

func main() {
	address := CreateAddress().
		SetStreet("Main Street").
		SetNumber(123).
		SetCity("New York").
		SetCountry("USA")

	fmt.Printf("address: %#v\n", address)
}

type Address struct {
	street  string
	number  int
	city    string
	country string
}

func CreateAddress() *Address {
	return &Address{}
}

func (a *Address) SetStreet(street string) *Address {
	a.street = street
	return a
}

func (a *Address) SetNumber(number int) *Address {
	a.number = number
	return a
}

func (a *Address) SetCity(city string) *Address {
	a.city = city
	return a
}

func (a *Address) SetCountry(country string) *Address {
	a.country = country
	return a
}
