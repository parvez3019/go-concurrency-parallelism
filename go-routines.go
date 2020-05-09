package main

import (
	"fmt"
	"runtime"
	"time"
)
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func RunRestaurant() {
	defer elapsed("page")()
	//runtime.GOMAXPROCS(1)
	jobs := make(chan int, 50)
	results := make(chan int, 50)

	go workers(jobs, results) //9.71276541s
	go workers(jobs, results) //5.887356294s
	go workers(jobs, results) //4.810860706s
	go workers(jobs, results) //4.463969722s
	go workers(jobs, results) //4.182831529s
	//go workers(jobs, results) //4.07890426s
	//go workers(jobs, results) //3.962344255s
	//go workers(jobs, results) //3.792776172s

	for i := 0; i < 50; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 50; j++ {
		fmt.Println(<-results)
	}

}

func workers(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func TestChannels() {
	const GOMAXPROCS = 4 //A
	runtime.GOMAXPROCS(GOMAXPROCS)

	c1 := make(chan string)
	c2 := make(chan string)

	go func(c1 chan string) {
		for char := 0; char < 26; char++ {
			c1 <- fmt.Sprintf("%c ", 'A'+char)
		}
		close(c1)
	}(c1)

	go func(c2 chan string) {
		for char := 0; char < 26; char++ {
			c2 <- fmt.Sprintf("%c ", 'a'+char)
		}
		close(c2)
	}(c2)

	for {
		select {
		case x, ok := <-c1:
			fmt.Println("ch1", x)
			if !ok {
				c1 = nil
			}
		case x, ok := <-c2:
			fmt.Println("ch2", x)
			if !ok {
				c2 = nil
			}
		}
		if c1 == nil && c2 == nil {
			break
		}
	}
}
