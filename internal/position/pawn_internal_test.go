package position

import (
	"testing"

	"github.com/samwestmoreland/chessengine/internal/board"
)

func TestPawnIsDoubled(t *testing.T) {
	sq1 := board.NewSquareOrPanic("a2")
	pawn1 := NewPawn(sq1, board.Black)

	sq2 := board.NewSquareOrPanic("a3")
	pawn2 := NewPawn(sq2, board.Black)

	pos := NewPosition(board.Black, []Piece{pawn1, pawn2})

	if !pawn1.IsDoubled(pos) || !pawn2.IsDoubled(pos) {
		t.Errorf("Pawns should be doubled")
	}
}
