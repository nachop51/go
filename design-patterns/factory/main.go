package main

import (
	prodcuts "factory/products"
	"fmt"
)

func main() {
	factory := prodcuts.Product{}

	products := factory.New()

	// fmt.Println(products)

	products.Name = "Product 1"

	fmt.Printf("Product: %+v\n", products)
}
