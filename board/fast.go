package board

type FastBoard struct {
	baseBoard *Board
}

func NewFastBoard(size int, population int) *FastBoard {
	board := NewBoard(size, population)
	fastBoard := FastBoard{board}

	return &fastBoard
}

func (board *FastBoard) Tick() {
	board.baseBoard.Tick()
}

func (board *FastBoard) String() string {
	return board.baseBoard.String()
}
