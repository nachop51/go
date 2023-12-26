package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func readDir(dirname string) []fs.FileInfo {

	if dirname == "" {
		dirname = "."
	}

	data, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	var files []fs.FileInfo

	for _, file := range data {

		fileInfo, err := file.Info()

		if err != nil {
			log.Fatal(err)
		}

		files = append(files, fileInfo)
	}

	return files
}

func updateFilesMap(filesMap map[string]time.Time) map[string]time.Time {

	if filesMap == nil {
		filesMap = make(map[string]time.Time)
	}

	var files = readDir(".")

	for _, file := range files {
		value, ok := filesMap[file.Name()]

		if ok && value.Equal(file.ModTime()) {
			continue
		} else if ok {
			fmt.Println("File modified:", file.Name())
		} else {
			fmt.Println("New file:", file.Name())
		}

		filesMap[file.Name()] = file.ModTime()
	}

	return filesMap
}

func main() {
	// var data string = readFile("data.txt")
	// data := readFile("test.txt")

	// fmt.Println(data)

	// for _, file := range files {
	// 	fmt.Println(file.IsDir(), file.Mode(), "'"+file.Name()+"'", "'"+fmt.Sprint(file.Size())+"'", file.ModTime())
	// }

	var filesMap = updateFilesMap(nil)

	for key, value := range filesMap {
		fmt.Println(key+":", value)
	}
	fmt.Printf("%p\n", filesMap)

	for {
		time.Sleep(1 * time.Second)
		filesMap = updateFilesMap(filesMap)
		fmt.Println(filesMap)
		fmt.Printf("%p\n", filesMap)
	}

}
