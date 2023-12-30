package main

import (
	"log"
	"os"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func readDir(dirname string) []os.FileInfo {
	// directories must be maps, and files must be slices

	data, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	var files []os.FileInfo

	for _, file := range data {

		fileInfo, err := file.Info()

		if err != nil {
			log.Fatal(err)
		}

		files = append(files, fileInfo)
	}

	return files
}

func updateFilesMap(filesMap map[string]os.FileInfo, dirname string) (map[string]os.FileInfo, bool) {
	if filesMap == nil {
		filesMap = make(map[string]os.FileInfo, 0)
	}

	var files = readDir(dirname)
	var updated bool = false

	for _, file := range files {
		if file.IsDir() {
			_, update := updateFilesMap(filesMap, dirname+"/"+file.Name())
			if update {
				updated = true
			}
			continue
		}
		if mapFile, ok := filesMap[dirname+"/"+file.Name()]; !ok {
			updated = true
			filesMap[dirname+"/"+file.Name()] = file
		} else {
			if mapFile.ModTime() != file.ModTime() {
				updated = true
				filesMap[dirname+"/"+file.Name()] = file
			}
		}
	}

	return filesMap, updated
}
