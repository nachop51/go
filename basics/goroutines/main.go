package main

import (
	"fmt"
	"net/http"
)

func demo() {
	var links []string = []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkStatus2(link, c)
	}

	// * Blocking call
	// & Waits for the channel to return a value
	// fmt.Println(<-c)

	// * Looping through the channel
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLinkStatus2(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- fmt.Sprint(link, " might be down!")
	} else {
		c <- fmt.Sprint(link, " is up!")
	}
}
