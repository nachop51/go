package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		log.Fatal(err)
	}

	// * resp.Body is a ReadCloser
	// * which means it implements the Reader and Closer interfaces
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Printf("%s", bs)

	// fmt.Printf("%+v", resp)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("The length of the byte slice is", len(bs), "bytes.")

	return len(bs), nil
}
