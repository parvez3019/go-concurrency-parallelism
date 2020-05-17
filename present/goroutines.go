package main

import (
	"fmt"
	"net/http"
	"time"
)
// START OMIT
func main() {
	sites := []string {"google", "facebook", "thoughtworks", "github"}

	start := time.Now()
	for _, site := range sites {
		go printResponse(site)
	}

	fmt.Println("Total time taken for the responses: ", time.Since(start))
}

func printResponse(site string) {	
	res, _ := http.Get("https://" + site + ".com")
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT
