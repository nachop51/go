package main

import "fmt"

func maps() {

	// var colors map[string]string
	// fmt.Println(colors) // prints out -> map[]

	// colors := map[string]string{
	// 	"red":   "#f00",
	// 	"green": "#0f0",
	// 	"blue":  "#00f",
	// }

	colors := make(map[string]string)

	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"

	// fmt.Println(colors["red"])

	// To delete a key-value pair from a map

	// delete(colors, "red")

	// fmt.Printf("%+v\n", colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
