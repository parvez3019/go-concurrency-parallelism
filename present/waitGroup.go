package main

import (
	"fmt"
	"net/http"
	"sync"
)
// START OMIT
func main() {
	sites := []string {"www.google.com", "www.facebook.com", "www.thoughtworks.com"}

	var wg sync.WaitGroup

	for _, site := range sites {
		wg.Add(1)
		go printResponse(site, &wg)
	}
	wg.Wait()
	fmt.Println("exit main")
}

func printResponse(site string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get("https://" + site)
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT
