package main

import (
	"fmt"
	"net/http"
	"time"
)

// 1. DATA MODELING
// We don't just want strings; we want data we can analyze later.
type Result struct {
	URL        string
	StatusCode int
	IsUp       bool
}

func main() {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://go.dev",
		"https://amazon.com",
	}

	
	jobs := make(chan string, len(websites))
	// results: A buffered channel to hold the finished reports
	results := make(chan Result, len(websites))

	
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	start := time.Now()

	
	for _, url := range websites {
		jobs <- url
	}
	close(jobs) // "No more jobs to add!"

	for i := 0; i < len(websites); i++ {
		result := <-results
		printResult(result)
	}

	fmt.Printf("\nTotal time taken: %s\n", time.Since(start))
}

func worker(id int, jobs <-chan string, results chan<- Result) {
	
	for url := range jobs {
		
		
		resp, err := http.Get(url)
		
		result := Result{URL: url, IsUp: true}
		
		if err != nil {
			result.IsUp = false
			result.StatusCode = 0
		} else {
			result.StatusCode = resp.StatusCode
		}

		results <- result
	}
}

// Helper to make printing cleaner
func printResult(r Result) {
	if r.IsUp {
		fmt.Printf("[%d] %s is UP\n", r.StatusCode, r.URL)
	} else {
		fmt.Printf("[DOWN] %s\n", r.URL)
	}
}