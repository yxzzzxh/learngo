package main

import (
	"fmt"
	"net/http"
)

type Msg struct {
	url     string
	isValid string // "FAILED" , "SUCCEED"
}

func main() {
	c := make(chan Msg)
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.naver.com",
		"https://www.airbnb.com",
		"https://www.facebook.com",
		"https://www.daum.com",
		"https://www.twitter.com",
		"https://www.tistory.com",
		"https://www.coupang.com",
		"https://www.nate.com",
	}

	for _, url := range urls {
		go IsValid(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

// Check a url
func IsValid(url string, c chan Msg) {
	fmt.Println("Checking " + url)
	resp, err := http.Get(url)
	msg := Msg{}
	msg.url = url
	if err != nil || resp.StatusCode >= 400 {
		msg.isValid = "FAILED"
	}
	msg.isValid = "SUCCEED"
	c <- msg
}
