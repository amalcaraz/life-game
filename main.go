package main

import (
	"fmt"
	"life/board"
	"life/utils"
	"time"
)

func main() {
	utils.Init()

	board := board.New(50, 500)
	board.Init()

	for {
		fmt.Print("\033[2J", board)
		time.Sleep(time.Millisecond * 200)
		board = board.Tick()
	}
}
