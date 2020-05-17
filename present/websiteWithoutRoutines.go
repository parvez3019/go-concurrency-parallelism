package main

import (
	"fmt"
	"net/http"
	"time"
)
// START OMIT
func main() {
	sites := []string {
		"google.com",
		"facebook.com",
		"thoughtworks.com",
		"github.com",
	}

	start := time.Now()

	for _, site := range sites {
		printResponse(site)
	}

	fmt.Println("Total time taken for the reponses: ", time.Since(start))
}

func printResponse(site string) {
	res, _ := http.Get("https://" + site)
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT


