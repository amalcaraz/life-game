package board

import "testing"

func BenchmarkFastBoard(b *testing.B) {
	board := NewFastBoard(1000, 500)

	for i := 0; i < b.N; i++ {
		board.Tick()
	}
}
