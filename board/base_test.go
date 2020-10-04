package board

import "testing"

func BenchmarkBoard(b *testing.B) {
	board := NewBoard(1000, 500)

	for i := 0; i < b.N; i++ {
		board.Tick()
	}
}
