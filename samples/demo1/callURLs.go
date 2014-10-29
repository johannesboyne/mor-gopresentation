package main

import (
	"fmt"
	"net/http"
)

func main() {
	response := make(chan string)
	urls := []string{
		"http://zweitag.de",
		"http://bleacherreport.com",
		"http://google.com",
	}

	for _, url := range urls {
		go callURL(url, response)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println("response from " + <-response)
	}
}

func callURL(url string, response_channel chan string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	response_channel <- res.Status + " " + url
}
