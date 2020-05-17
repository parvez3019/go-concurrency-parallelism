package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {

	var Ball int
	table := make(chan int)

	go player(table, "Player 1")
	go player(table, "Player 2")

	table <- Ball
	time.Sleep(10 * time.Second)
	<-table
}

func player(table chan int, playerName string) {
	for {
		ball := <-table
		fmt.Printf("%s taking ball from %d table\n", playerName, ball)
		ball++
		time.Sleep(1 * time.Second)
		fmt.Printf("%s putting ball %d on table\n", playerName, ball)
		table <- ball
	}
}
// END OMIT