package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	sites := []string{
		"www.google.com",
		"www.facebook.com",
		"www.thoughtworks.com",
		"www.github.com",
	}

	var wg sync.WaitGroup

	result := make(chan string, 4)

	start := time.Now()

	for _, site := range sites {
		wg.Add(1)
		go sendResponse(site, &wg, result)
	}
	wg.Wait()
	close(result)

	for r := range result {
		fmt.Println(r)
	}
	fmt.Println("Total time taken for the responses: ", time.Since(start))
}

// START OMIT
func sendResponse(site string, wg *sync.WaitGroup, result chan string) {
	defer wg.Done()
	res, _ := http.Get("https://" + site)
	result <- fmt.Sprintf("%s response -> %d", site, res.StatusCode)
}
// END OMIT
