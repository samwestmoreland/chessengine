package position

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/pieces"
)

func TestGetPositionFromFEN(t *testing.T) {
	fen, err := ParseFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	if err != nil {
		t.Errorf("Error in ParseFEN: %s", err)
	}
	pos := getPositionFromFEN(fen)
	if pos == nil {
		t.Error("Error in GetPositionFromFEN")
	}

	// Check that there is a pawn on e4
	e4, err := board.ParseSquare("e4")
	if err != nil {
		t.Errorf("Error in ParseSquare: %s", err)
	}
	piece := pos.White[e4]
	if piece == nil {
		t.Error("Error in GetPositionFromFEN")
	}
	if piece.Type() != pieces.PawnType {
		t.Error("Error in GetPositionFromFEN")
	}
}
