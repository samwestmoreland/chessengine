package position

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
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
	p := pos.White[e4]
	if p == nil {
		t.Error("Error in GetPositionFromFEN")
	}
	if p.Type() != piece.PawnType {
		t.Error("Error in GetPositionFromFEN")
	}
}

func TestGetMovesForKing(t *testing.T) {
	sqStr := "e4"
	sq, err := board.NewSquare(sqStr)
	if err != nil {
		t.Fatalf("Failed to create square %s: %v", sqStr, err)
	}
	whiteKing := NewKing(sq, board.White)
	sqStrBlack := "h8"
	sqBlack, err := board.NewSquare(sqStrBlack)
	if err != nil {
		t.Errorf("Failed to create square %s", sqStr)
	}
	blackKing := NewKing(sqBlack, board.Black)

	pos := NewPosition(board.White, []Piece{whiteKing, blackKing})

	moves, err := whiteKing.GetMoves(*sq, pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	fmt.Printf("moves: %v", moves)
}
