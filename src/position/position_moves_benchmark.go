package position

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
)

func BenchmarkGetAllMovesConcurrent(b *testing.B) {
	e4 := board.NewSquareOrPanic("e4")
	blackKing := NewKing(e4, board.Black)

	g3 := board.NewSquareOrPanic("g3")
	blackBishop := NewBishop(g3, board.Black)

	a8 := board.NewSquareOrPanic("a8")
	blackRook := NewRook(a8, board.Black)

	c1 := board.NewSquareOrPanic("c1")
	blackQueen := NewQueen(c1, board.Black)

	g4 := board.NewSquareOrPanic("g4")
	blackBishop2 := NewBishop(g4, board.Black)

	pos := NewPosition(board.Black, []Piece{blackKing, blackBishop, blackRook, blackQueen, blackBishop2})

	for i := 0; i < b.N; i++ {
		_, err := pos.GetAllMovesConcurrent(board.Black)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetAllMovesSerial(b *testing.B) {
	e4 := board.NewSquareOrPanic("e4")
	blackKing := NewKing(e4, board.Black)

	g3 := board.NewSquareOrPanic("g3")
	blackBishop := NewBishop(g3, board.Black)

	a8 := board.NewSquareOrPanic("a8")
	blackRook := NewRook(a8, board.Black)

	c1 := board.NewSquareOrPanic("c1")
	blackQueen := NewQueen(c1, board.Black)

	g4 := board.NewSquareOrPanic("g4")
	blackBishop2 := NewBishop(g4, board.Black)

	pos := NewPosition(board.Black, []Piece{blackKing, blackBishop, blackRook, blackQueen, blackBishop2})

	for i := 0; i < b.N; i++ {
		_, err := pos.GetAllMovesSerial(board.Black)
		if err != nil {
			b.Error(err)
		}
	}
}
