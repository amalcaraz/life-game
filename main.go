package main

import (
	"fmt"
	"life/board"
	"time"
)

func main() {
	update := make(chan string)
	go run(update)

	for {
		fmt.Print(<-update)
	}
}

func run(update chan string) {
	game := board.NewFastBoard(55, 1000)

	for {
		game.Tick()
		update <- "\033[3J\033[;H" + game.String()
		time.Sleep(time.Millisecond * 100)
	}
}
