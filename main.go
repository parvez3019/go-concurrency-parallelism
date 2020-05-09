package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	// a := make([]int, 10)
	// variable initialization
	// slices
	// make
	// for loops
	// function calling -> function can return multiple values
	// GOMAXPROCS
	// error
	// main -> go routine -> simple will be to add time.second
	// wait groups

	//listOfSites := []string{
	//	"www.google.com",
	//	"www.facebook.com",
	//	"www.amazon.com",
	//	"www.stackoverflow.com",
	//	"www.linkedin.com",
	//	"www.github.com",
	//	"jigsaw.thoughtworks.net",
	//}

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	start := time.Now()

	a := []int{5, 10, 25, 35, 44, 45, 1}
	for _, site := range a {
		wg.Add(1)
		//go get(&wg, site)
		go fib(&wg, site)
	}
	wg.Wait()

	fmt.Printf("time elapsed %v \n", time.Since(start))
	fmt.Println("This is an awesome geek night!")
	fmt.Println(runtime.NumCPU())
}

func fib(s *sync.WaitGroup, site int) {
	fmt.Println(getfib(site))
	s.Done()
}

func getfib(site int) int {
	if site < 2 {
		return site
	}
	return getfib(site-1) + getfib(site-2)
}

func get(wg *sync.WaitGroup, site string) {
	res, _ := http.Get("https://" + site)
	fmt.Printf("%s response -> %d \n", site, res.StatusCode)
	wg.Done()
}
