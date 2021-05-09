package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
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
		IsValid(url)
	}
}

// Check a url
func IsValid(url string) error {
	fmt.Println("Checking " + url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errors.New(url + "is not valid")
	}
	return nil
}
