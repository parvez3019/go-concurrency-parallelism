package main

import (
	"fmt"
	"net/http"
	"time"
)
// START OMIT
func main() {
	sites := [4]string {"google", "facebook", "youtube", "github"}

	start := time.Now()
	for _, site := range sites {
		go printResponse(site)
	}

	time.Sleep(time.Second * 3)
	fmt.Println("Total time taken for the responses: ", time.Since(start))
}

func printResponse(site string) {
	res, _ := http.Get("https://www." + site + ".com")
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
}
// END OMIT
