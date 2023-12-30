package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// import (
// "fmt"
// "os"
// "os/exec"
// "time"
// )

func startNewProcess(currentProcess *os.Process) *os.Process {
	var err error

	if currentProcess != nil {
		fmt.Printf("Killing process %d\n", currentProcess.Pid)
		currentProcess.Release()
		currentProcess.Kill()
		currentProcess.Wait()
	}

	fmt.Println("Starting new process")
	currentProcess, err = os.StartProcess("/usr/local/go/bin/go", []string{"go", "run", "."}, &os.ProcAttr{
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	})

	fmt.Println("currentProcess:", currentProcess.Pid)

	if err != nil {
		log.Fatal(err)
	}

	return currentProcess
}

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
	var currentProcess *os.Process = nil

	// fmt.Println(filesMap, updated)

	for {
		if updated {
			currentProcess = startNewProcess(currentProcess)

			fmt.Println("currentProcess:", currentProcess)

			fmt.Println("Done!")
		}

		time.Sleep(1 * time.Second)
		filesMap, updated = updateFilesMap(filesMap, ".")
	}
}
