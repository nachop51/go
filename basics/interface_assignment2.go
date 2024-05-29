package main

import (
	"io"
	"log"
	"os"
)

func interface_assignment2(args []string) {
	var filename string

	if len(args) < 2 {
		log.Fatal("Please provide a filename")
	}

	filename = args[1]

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	io.Copy(os.Stdout, file)

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Done reading file")
}
