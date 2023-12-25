package main

import (
	"fmt"
)

type Person struct {
	name string
	age  float64
}

func main() {
	var x = []int{1, 2, 3, 4, 5}

	y := x[1:3]

	x[4] = 20

	ls := make([]int, 0, 20)

	ls = append(ls, 1, 2, 3, 4, 5)

	ls = append(ls, x...)
	ls = append(ls, x...)
	ls = append(ls, x...)
	ls = append(ls, x...)

	fmt.Println("Hello, World!", x, y, ls)

	var str = "Hello, World!"

	fmt.Println(str[0:5])
	fmt.Println(str[7:12])

	var p = Person{"John", 20.11}

	fmt.Println(p.name, p.age)

	var p2 = Person{age: 24.52, name: "Another John"}

	fmt.Println(p2.name, p2.age)

	const PI float64 = 3.14159265358979323846264338327950288419716939937510582097494459
	const PI2 = 3.14159265358979323846264338327950288419716939937510582097494459

	fmt.Printf("PI = %f\n", PI)
	fmt.Println(PI2)

	// enums
	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
		TB
	)

	const (
		_   = iota
		KB2 = 1 << (10 * iota)
		MB2
		GB2
		TB2
	)

	fmt.Println(KB, MB, GB, TB)
	fmt.Println(KB2, MB2, GB2, TB2)
}
