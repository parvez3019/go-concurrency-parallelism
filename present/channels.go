package main

import (
	"fmt"
	"net/http"
	"sync"
)


func main() {
	sites := []string{
		"google.com",
		"facebook.com",
		"thoughtworks.com",
		"github.com",
	}

	var wg sync.WaitGroup
	// START OMIT
	result := make(chan string, 4)
	for _, site := range sites {
		wg.Add(1)
		go printResponse(site, &wg, result)
	}
	wg.Wait()
	close(result)

	for r := range result {
		fmt.Println(r)
	}
	// END OMIT
}

func printResponse(site string, wg *sync.WaitGroup, result chan string) {
	defer wg.Done()
	res, _ := http.Get("https://" + site)
	result <- fmt.Sprintf("%s response -> %d \n", site, res.StatusCode)
}

