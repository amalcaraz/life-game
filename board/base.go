package board

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	cells      *[][]bool
	size       int
	population int
	generation int
}

func NewBoard(size int, population int) *Board {
	cells := newCells(size)

	board := Board{cells, size, population, 0}
	board.init()

	return &board
}

func newCells(size int) *[][]bool {
	cells := make([][]bool, size)

	for i := range cells {
		cells[i] = make([]bool, size)
	}

	return &cells
}

func (board *Board) init() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	pop := board.population
	cells := (*board.cells)

	for pop > 0 {
		i := r.Intn(board.size)
		j := r.Intn(board.size)

		cells[i][j] = true

		pop--
	}
}

func (board *Board) Tick() {
	cells := (*board.cells)
	newCells := (*newCells(board.size))
	newPopulation := 0

	for i := range cells {
		for j := range cells[i] {
			cell := board.calculateNextCell(i, j)
			newCells[i][j] = cell

			if cell {
				newPopulation++
			}
		}
	}

	board.cells = &newCells
	board.generation++
	board.population = newPopulation
}

func (board *Board) calculateNextCell(x int, y int) bool {
	count := 0
	cells := (*board.cells)

	for offsetX := -1; offsetX <= 1; offsetX++ {
		for offsetY := -1; offsetY <= 1; offsetY++ {

			x1 := x + offsetX
			y1 := y + offsetY

			if (x1 >= 0 && x1 < board.size) &&
				(y1 >= 0 && y1 < board.size) &&
				(x1 != x || y1 != y) {

				if cells[x1][y1] {
					count++
				}
			}
		}
	}

	cell := cells[x][y]

	if cell && (count < 2 || count > 3) {
		return false
	}

	if !cell && count == 3 {
		return true
	}

	return cell
}

func (board *Board) String() string {
	str := ""
	cells := (*board.cells)

	for i := range cells {
		for j := range cells {
			char := "◾️"

			if cells[i][j] == true {
				char = "◽️"
			}

			str += char
		}
		str += "\n"
	}

	str += fmt.Sprintf("Population: %d \n", board.population)
	str += fmt.Sprintf("Generation: %d \n", board.generation)

	return str
}
