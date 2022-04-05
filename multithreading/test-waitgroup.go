package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com",
		"http://www.google.com",
		"http://www.golang.org",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(url, resp.Status)
			} else {
				fmt.Println(url, " err")
			}
		}(url)
	}
	wg.Wait()
	fmt.Println("main end....")
}
