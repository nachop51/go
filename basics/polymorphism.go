package main

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Company string
}

func (p Person) PrintName() {
	println(p.Name)
}

func test() {
	emp := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
	}

	emp.PrintName()
	emp.Person.PrintName()
}
