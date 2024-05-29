package main

import "os"

func main() {
	// var i interface{} = os.Args
	// var i interface{} = "Hello"
	// var i interface{} = 42
	// var i interface{} = 3.14159
	// var i interface{} = true
	// var i interface{} = struct{ X int }{42}
	// var i interface{} = []int{1, 2, 3}
	// var i interface{} = map[string]int{"one": 1, "two": 2, "three": 3}
	// var i interface{} = make(chan int)
	// var i interface{} = func() { println("Hello, world!") }
	// var i interface{} = nil

	// switch v := i.(type) {
	// case []string:
	// 	println("[]string:", v)
	// case string:
	// 	println("string:", v)
	// case int:
	// 	println("int:", v)
	// case float64:
	// 	println("float64:", v)
	// case bool:
	// 	println("bool:", v)
	// case struct{ X int }:
	// 	println("struct{ X int }:", v)
	// case []int:
	// 	println("[]int:", v)
	// case map[string]int:
	// 	println("map[string]int:", v)
	// case chan int:
	// 	println("chan int:", v)
	// case func():
	// 	println("func():", v)
	// case nil:
	// 	println("nil")
	// default:
	// 	println("unknown")
	// }
	// maps()
	// structs()
	// interfaces()
	// interface_assignment()
	interface_assignment2(os.Args)
}
