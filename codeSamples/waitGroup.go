package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)
// START OMIT
func main() {
	sites := [4]string {"google", "facebook", "youtube", "github"}

	var wg sync.WaitGroup

	start := time.Now()
	for _, site := range sites {
		wg.Add(1)
		go printResponse(site, &wg)
	}
	wg.Wait()
	fmt.Println("Total time taken for the responses: ", time.Since(start))
}

func printResponse(site string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get("https://www." + site + ".com")
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT
