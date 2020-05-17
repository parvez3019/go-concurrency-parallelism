package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	var Ball int
	table := make(chan int)
	result := make(chan string)

	go player(table, result, "Player 1")
	go player(table, result, "Player 2")

	table <- Ball

	for {
		select {
		case a := <-result :
			fmt.Println(a)
		}
	}

}

func player(table chan int, result chan string, player string) {
	for {
		ball := <-table
		ball++
		time.Sleep(1 * time.Second)
		table <- ball
		result <- fmt.Sprintf("%s turn number : %d", player, ball)
	}
}

// END OMIT
