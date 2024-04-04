package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // Add  "go" before function call for new go routine
	}

	// fmt.Println(<-c) // Channel values will be printed, also a blocking line of code

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Blocking call: For every url, we wait for the response to comeback and then proceed with further links

	if err != nil {
		fmt.Println(link, "loading error")
		// c <- "Might be down" // Send value to channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yup, It's up" // Send value to channel
	c <- link
}
