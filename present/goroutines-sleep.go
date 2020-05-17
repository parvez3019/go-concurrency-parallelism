package main

import (
	"fmt"
	"net/http"
	"time"
)
// START OMIT
func main() {
	sites := []string {
		"www.google.com",
		"www.facebook.com",
		"www.thoughtworks.com",
		"github.com",
	}

	start := time.Now()

	for _, site := range sites {
		go printResponse(site)
	}

	fmt.Scanln()
	fmt.Println("Total time taken for the responses: ", time.Since(start))
}

func printResponse(site string) {
	res, _ := http.Get("https://" + site)
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT