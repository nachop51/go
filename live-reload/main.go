package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// var data string = readFile("data.txt")
	// data := readFile("test.txt")
	//
	// fmt.Println(data)
	//
	// for _, file := range files {
	// 	fmt.Println(file.IsDir(), file.Mode(), "'"+file.Name()+"'", "'"+fmt.Sprint(file.Size())+"'", file.ModTime())
	// }
	//
	// var filesMap = make(map[string]time.Time)
	// or just
	// filesMap := map[string]time.Time{}
	//
	// for key, value := range filesMap {
	// 	fmt.Println(key+":", value)
	// }
	// fmt.Printf("%p\n", filesMap)

	var filesMap, updated = updateFilesMap(nil, ".")

	// fmt.Println(filesMap, updated)

	for {
		if updated {
			fmt.Println("Initializing...")

			cmd := exec.Command("go", "run", ".")

			cmd.Stdout = os.Stdout

			err := cmd.Run()

			if err != nil {
				fmt.Println(err)
			}

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Done!")
		}

		time.Sleep(1 * time.Second)
		filesMap, updated = updateFilesMap(filesMap, ".")
	}
}
