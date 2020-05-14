package main

import (
	"fmt"
	"net/http"
	"time"
)
// START OMIT
func main() {
	sites := []string {"www.google.com", "www.facebook.com", "www.thoughtworks.com"}
	for _, site := range sites {
       		go printResponse(site)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("exit main")
}

func printResponse(site string) {
	res, _ := http.Get("https://" + site)
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT
