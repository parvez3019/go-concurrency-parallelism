package main

import (
	"fmt"
	"sync"
	"time"

)

func main() {
	// START OMIT
	var wg sync.WaitGroup
	start := time.Now()

	a := [4]int{38, 39, 40, 41}
	for _, num := range a {
		wg.Add(1)
		go fib(&wg, num)
	}
	wg.Wait()
	fmt.Printf("time elapsed %v \n", time.Since(start))
	// END OMIT
}

func fib(s *sync.WaitGroup, num int) {
	fmt.Println(getfib(num))
	s.Done()
}

func getfib(site int) int {
	if site < 2 {
		return site
	}
	return getfib(site-1) + getfib(site-2)
}
