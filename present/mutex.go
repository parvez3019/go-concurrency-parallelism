package main

import (
	"fmt"
	"net/http"
	"sync"
)
// START OMIT
var mutex sync.Mutex

func printResponse(site string, wg *sync.WaitGroup, result map[string]int) {
        defer wg.Done()
        res, _ := http.Get("https://" + site)
        mutex.Lock()
        result[site] = res.StatusCode
        mutex.Unlock()
}
// END OMIT
func main() {
	sites := []string {"www.google.com", "www.facebook.com", "www.thoughtworks.com"}

	var wg sync.WaitGroup
	result := map[string]int{}
	for _, site := range sites {
		wg.Add(1)
		go printResponse(site, &wg, result)
	}
	wg.Wait()
	
	for k,v := range result {
		fmt.Printf("%s -> %d\n", k,v)
	}
	fmt.Println("exit main")
}
