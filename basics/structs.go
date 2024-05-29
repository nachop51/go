package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
	// contact   contactInfo
}

func structs() {
	// var p1 person = person{
	// 	firstName: "James",
	// 	lastName:  "Bond",
	// 	age:       32,
	// }

	// p2 := person{firstName: "Alex", lastName: "Anderson", age: 22}

	// fmt.Println(p1)
	// fmt.Println(p2)

	// fmt.Printf("%T\n", p1)
	// fmt.Printf("%+v\n", p2)

	// p1.age = 33

	// fmt.Printf("%v\n", p1)

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	age:       33,
	// 	contactInfo: contactInfo{
	// 		email:   "hewnupar@huzoace.se",
	// 		zipCode: 94000,
	// 	},
	// }

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jim.print()

	jim := &person{
		firstName: "Jim",
		lastName:  "Party",
		age:       33,
		contactInfo: contactInfo{
			email:   "hewnupar@huzoace.se",
			zipCode: 94000,
		},
	}
	jim.updateName("Jimmy")

	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v\n", p)
}
