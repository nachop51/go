package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var links []string = []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkStatus(link, c)
	}

	// * Looping through the channel
	// * Infinite loop
	// l is the value from the channel
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLinkStatus(link, c)
		}(l)
	}
}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
	} else {
		fmt.Println(link, " is up!")
	}

	c <- link
}
