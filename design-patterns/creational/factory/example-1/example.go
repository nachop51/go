package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Magazine 1", 100, "Publisher 1")
	mag2, _ := newPublication("magazine", "Magazine 2", 200, "Publisher 2")

	news1, _ := newPublication("newspaper", "Newspaper 1", 10, "Publisher 1")
	news2, _ := newPublication("newspaper", "Newspaper 2", 20, "Publisher 2")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub iPublication) {
	fmt.Printf("--------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("--------------------\n")
}
