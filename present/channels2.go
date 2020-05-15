package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	sites := []string{"www.google.com", "www.facebook.com", "www.thoughtworks.com"}

	var wg sync.WaitGroup

	result := make(chan string, 3)
	for _, site := range sites {
		wg.Add(1)
		go printResponse(site, &wg, result)
	}
	wg.Wait()
	close(result)

	for r := range result {
		fmt.Println(r)
	}
	fmt.Println("main exit")
}

// START OMIT
func printResponse(site string, wg *sync.WaitGroup, result chan string) {
	defer wg.Done()
	res, _ := http.Get("https://" + site)
	result <- fmt.Sprintf("%s response -> %d", site, res.StatusCode)
}
// END OMIT
