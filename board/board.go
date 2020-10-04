package board

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	cells      [][]bool
	size       int
	population int
	generation int
}

func New(size int, population int) *Board {
	cells := make([][]bool, size)

	for i := range cells {
		cells[i] = make([]bool, size)
	}

	board := Board{cells, size, population, 0}

	return &board
}

func (board *Board) Init() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	pop := board.population

	for pop > 0 {
		i := r.Intn(board.size)
		j := r.Intn(board.size)

		board.cells[i][j] = true

		pop--
	}
}

func (board *Board) Tick() *Board {
	newBoard := New(board.size, 0)

	for i := range board.cells {
		for j := range board.cells[i] {
			cell := board.calculateCell(i, j)
			newBoard.cells[i][j] = cell

			if cell {
				newBoard.population++
			}
		}
	}

	newBoard.generation = board.generation + 1
	return newBoard
}

func (board *Board) calculateCell(i int, j int) bool {
	count := 0

	for h := -1; h < 2; h++ {
		for v := -1; v < 2; v++ {

			x := i + h
			y := j + v

			if (x > 0 && x < board.size) &&
				(y > 0 && y < board.size) &&
				(x != i || y != j) {

				localCell := board.cells[x][y]

				if localCell {
					count++
				}
			}
		}
	}

	cell := board.cells[i][j]

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

	for i := range board.cells {
		for j := range board.cells {
			char := "⬜️"

			if board.cells[i][j] == true {
				char = "⬛"
			}

			str += char
		}
		str += "\n"
	}

	str += fmt.Sprintf("Population: %d \n", board.population)
	str += fmt.Sprintf("Generation: %d \n", board.generation)

	return str
}
