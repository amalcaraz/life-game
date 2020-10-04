package main

import (
	"fmt"
	"life/board"
	"time"
)

func main() {

	board := board.NewBoard(50, 500)
	run(board)
}

func run(board board.BoardI) {
	for {
		// Note: Clear the terminal bufer
		// fmt.Print("\033[2J\033[3J\033[;H", board)
		fmt.Print("\033[3J\033[;H", board)

		time.Sleep(time.Millisecond * 100)
		board.Tick()
	}
}
