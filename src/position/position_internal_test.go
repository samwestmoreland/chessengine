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
		t.Fatalf("Error in ParseFEN: %s", err)
	}

	pos := getPositionFromFEN(fen)
	if pos == nil {
		t.Fatal("Error in GetPositionFromFEN")
	}

	// Print the position
	t.Log(pos.String())

	// Check that there is a pawn on e4
	e4, err := board.NewSquare("e4")
	if err != nil {
		t.Fatalf("Error initialising square: %s", err)
	}

	p := pos.White[e4]
	if p == nil {
		t.Fatal("got nil piece but was expecting a pawn")
	}

	if p.Type() != piece.PawnType {
		t.Fatal("Error in GetPositionFromFEN")
	}
}

func TestGetMovesForKingOnEmptyBoard(t *testing.T) {
	sqStr := "e4"

	square, err := board.NewSquare(sqStr)
	if err != nil {
		t.Fatalf("Failed to create square %s: %v", sqStr, err)
	}

	whiteKing := NewKing(square, board.White)
	pos := NewPosition(board.White, []Piece{whiteKing})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	d5, _ := board.NewSquare("d5")
	e5, _ := board.NewSquare("e5")
	f5, _ := board.NewSquare("f5")
	d4, _ := board.NewSquare("d4")
	f4, _ := board.NewSquare("f4")
	d3, _ := board.NewSquare("d3")
	e3, _ := board.NewSquare("e3")
	f3, _ := board.NewSquare("f3")

	expectedMoves := []moves.Move{
		{From: square, To: d5, PieceType: piece.KingType},
		{From: square, To: e5, PieceType: piece.KingType},
		{From: square, To: f5, PieceType: piece.KingType},
		{From: square, To: d4, PieceType: piece.KingType},
		{From: square, To: f4, PieceType: piece.KingType},
		{From: square, To: d3, PieceType: piece.KingType},
		{From: square, To: e3, PieceType: piece.KingType},
		{From: square, To: f3, PieceType: piece.KingType},
	}

	if len(mov) != len(expectedMoves) {
		t.Errorf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	// Check that the moves are the same, but don't care about order
	for _, expectedMove := range expectedMoves {
		found := false

		for _, move := range mov {
			if expectedMove.Equals(move) {
				found = true

				break
			}
		}

		if !found {
			t.Errorf("Expected move %v not found", expectedMove)
		}
	}
}

func TestGetMovesForKingOnEmptyBoardInCorner(t *testing.T) {
	sqStr := "a1"

	square, err := board.NewSquare(sqStr)
	if err != nil {
		t.Fatalf("Failed to create square %s: %v", sqStr, err)
	}

	whiteKing := NewKing(square, board.White)

	pos := NewPosition(board.White, []Piece{whiteKing})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	b1, _ := board.NewSquare("b1")
	a2, _ := board.NewSquare("a2")
	b2, _ := board.NewSquare("b2")

	expectedMoves := []moves.Move{
		{From: square, To: b1, PieceType: piece.KingType},
		{From: square, To: a2, PieceType: piece.KingType},
		{From: square, To: b2, PieceType: piece.KingType},
	}

	if len(mov) != len(expectedMoves) {
		t.Errorf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	// Check that the moves are the same, but don't care about order
	for _, expectedMove := range expectedMoves {
		found := false

		for _, move := range mov {
			if expectedMove.Equals(move) {
				found = true

				break
			}
		}

		if !found {
			t.Errorf("Expected move %v not found", expectedMove)
		}
	}
}