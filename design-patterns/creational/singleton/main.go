package main

import "fmt"

func main() {
	// log := GetLoggerInstance()

	// log.SetLogLevel(1)
	// log.Log("Hello, world!")

	// log = GetLoggerInstance()
	// log.SetLogLevel(2)
	// log.Log("Hello again!")

	// log = GetLoggerInstance()
	// log.SetLogLevel(3)
	// log.Log("Hello one more time!")

	for i := 0; i < 10; i++ {
		go GetLoggerInstance()
	}

	fmt.Scanln()
}
