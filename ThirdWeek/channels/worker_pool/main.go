package main

import (
	"io/ioutil"
	"net/http"
)

func getSiteContent(id int, url <-chan string, result chan<- string) {
	for v := range url {
		resp, _ := http.Get(v)
		defer resp.Body.Close()
		respBodyBytes, _ := ioutil.ReadAll(resp.Body)
		result <- string(respBodyBytes)
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://gmail.com",
		"https://youtube.com",
		"https://twitter.com",
	}

	url := make(chan string, len(urls))
	result := make(chan string, len(urls))

	for i := 0; i < 3; i++ {
		go getSiteContent(i, url, result)
	}

	for i := 0; i < len(urls); i++ {
		url <- urls[i]
	}

	close(url)

	for i := 0; i < len(urls); i++ {
		<-result
	}

}
