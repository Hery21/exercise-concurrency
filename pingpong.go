package main

import (
	"fmt"
)

type ball struct {
	hits int
}

func player(name string, table chan *ball) {
	for {
		ball := <-table // player grabs the ball
		ball.hits++     // hits the ball
		fmt.Println(name, "hits", ball.hits)
		table <- ball // pass the ball
	}
}

func main() {

	table := make(chan *ball)

	go player("player A", table)
	go player("player B", table)
	go player("player C", table)
	go player("player D", table)

	// before referee
	//table <- new(ball)
	//
	//time.Sleep(1 * time.Millisecond)
	//defer close(playerChannel)

	//var playerChannel = make(chan string, 1)
	//
	//playerA := "ping pong"
	//var playerB string
	//
	//for {
	//	time.Sleep(3 * time.Second)
	//	playerChannel <- playerA
	//	playerB = <-playerChannel
	//	playerA = ""
	//	fmt.Printf("%v", playerB)
	//}
}
