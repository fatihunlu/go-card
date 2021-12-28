package main

import (
	"fmt"
	"net/http"
	"time"
)

func getStatuses() {
	links := []string{
		"http://teamwork.com",
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 2)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down ..")
		c <- link
		return
	}

	fmt.Println(link, " is alive")
	c <- link
}
