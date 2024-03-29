package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)
// START OMIT
var mutex sync.Mutex

func printResponse(site string, wg *sync.WaitGroup, result *[]string) {
        defer wg.Done()
        res, _ := http.Get("https://" + site)
        mutex.Lock()
		*result = append(*result, fmt.Sprintf("%s -> %d", site, res.StatusCode))
        mutex.Unlock()
}
// END OMIT
func main() {
	sites := []string {"www.google.com", "www.facebook.com", "www.youtube.com", "www.github.com"}

	var wg sync.WaitGroup

	result := []string{}

	start := time.Now()
	for _, site := range sites {
		wg.Add(1)
		go printResponse(site, &wg, &result)
	}
	wg.Wait()
	
	for _,v := range result {
		fmt.Println(v)
	}

	fmt.Println("Total time taken for the responses: ", time.Since(start))
}
