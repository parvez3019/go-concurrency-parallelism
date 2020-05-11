package main

import (
	"fmt"
	"net/http"
)
// START OMIT
func main() {
	sites := []string {"www.google.com", "www.facebook.com", "www.thoughtworks.com"}

	for _, site := range sites {
		printResponse(site)
	}
}

func printResponse(site string) {
	res, _ := http.Get("https://" + site)
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT


