package board

import (
	"math"
	"sync"
)

type FastBoard struct {
	*Board
}

func NewFastBoard(size int, population int) *FastBoard {
	board := NewBoard(size, population)
	fastBoard := FastBoard{board}

	return &fastBoard
}

func (board *FastBoard) Tick() {
	newCells := newCells(board.size)
	newPopulation := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go board.calculateNextCells(0, 0, board.size, board.size, newCells, &newPopulation, &wg)
	wg.Wait()

	board.cells = newCells
	board.population = newPopulation
	board.generation++
}

func (board *FastBoard) calculateNextCells(x, y, xLen, yLen int, newCells *[][]bool, population *int, wg *sync.WaitGroup) {
	defer wg.Done()

	if xLen == 1 && yLen == 1 {

		cell := board.calculateNextCell(x, y)
		(*newCells)[x][y] = cell

		if cell {
			(*population)++
		}

		return
	}

	x0 := x
	x0Len := int(math.Ceil(float64(xLen) / 2))
	x1 := x0 + x0Len
	x1Len := xLen - x0Len

	y0 := y
	y0Len := int(math.Ceil(float64(yLen) / 2))
	y1 := y0 + y0Len
	y1Len := yLen - y0Len

	wg.Add(1)
	go board.calculateNextCells(x0, y0, x0Len, y0Len, newCells, population, wg)

	if x1Len > 0 {
		wg.Add(1)
		go board.calculateNextCells(x1, y0, x1Len, y0Len, newCells, population, wg)
	}

	if y1Len > 0 {
		wg.Add(1)
		go board.calculateNextCells(x0, y1, x0Len, y1Len, newCells, population, wg)
	}

	if x1Len > 0 && y1Len > 0 {
		wg.Add(1)
		go board.calculateNextCells(x1, y1, x1Len, y1Len, newCells, population, wg)
	}
}

func (board *FastBoard) String() string {
	return board.Board.String()
}
