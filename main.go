package main

import (
	"fmt"
	"net/http"
	"sync" // We need this back!
	"time"
)

func main() {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://go.dev",
		"https://amazon.com",
	}

	c := make(chan string)
	var wg sync.WaitGroup 

	start := time.Now()

	for _, site := range websites {
		wg.Add(1)
		go checkWebsite(site, c, &wg)
	}

	
	go func() {
		wg.Wait()
		close(c) 
	}()

	
	for msg := range c {
		fmt.Println(msg)
	}

	fmt.Printf("\nTotal time taken: %s\n", time.Since(start))
}

func checkWebsite(url string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done() 

	resp, err := http.Get(url)
	
	if err != nil {
		c <- fmt.Sprintf("[DOWN] %s", url)
		return
	}
	
	c <- fmt.Sprintf("[%d] %s is UP", resp.StatusCode, url)
}


