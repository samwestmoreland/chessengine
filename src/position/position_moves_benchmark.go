package position

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
)

func BenchmarkGetAllWhiteMoves(b *testing.B) {
	e4, _ := board.NewSquare("e4")
	g3, _ := board.NewSquare("g3")
	whiteKing := NewKing(e4, board.White)
	whiteBishop := NewBishop(g3, board.White)
	pos := NewPosition(board.White, []Piece{whiteKing, whiteBishop})

	for i := 0; i < b.N; i++ {
		_, err := pos.GetAllWhiteMoves()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetAllBlackMoves(b *testing.B) {
	e4, _ := board.NewSquare("e4")
	g3, _ := board.NewSquare("g3")
	blackKing := NewKing(e4, board.Black)
	blackBishop := NewBishop(g3, board.Black)
	pos := NewPosition(board.Black, []Piece{blackKing, blackBishop})

	for i := 0; i < b.N; i++ {
		_, err := pos.GetAllBlackMoves()
		if err != nil {
			b.Error(err)
		}
	}
}
