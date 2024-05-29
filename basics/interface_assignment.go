package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func interface_assignment() {
	sq := square{sideLength: 4}
	tr := triangle{height: 3, base: 5}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}
