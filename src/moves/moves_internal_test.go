package moves

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/piece"
)

func TestNewMove(t *testing.T) {
	fromSquare, _ := board.NewSquare("a1")
	toSquare, _ := board.NewSquare("a2")

	move := NewMove(fromSquare, toSquare, piece.QueenType)
	if move.From != fromSquare {
		t.Errorf("Expected %v, got %v", fromSquare, move.From)
	}

	if move.To != toSquare {
		t.Errorf("Expected %v, got %v", toSquare, move.To)
	}

	if move.PieceType != piece.QueenType {
		t.Errorf("Expected %v, got %v", piece.QueenType, move.PieceType)
	}
}

func TestMoveStringRepresentation(t *testing.T) {
	fromSquare, _ := board.NewSquare("a1")
	toSquare, _ := board.NewSquare("a2")

	move := NewMove(fromSquare, toSquare, piece.QueenType)

	expected := "Queen: a1 -> a2"
	if move.String() != expected {
		t.Errorf("Expected %v, got %v", expected, move.String())
	}
}

func TestMoveEquality(t *testing.T) {
	fromSquare, _ := board.NewSquare("a1")
	toSquare, _ := board.NewSquare("a8")

	move1 := NewMove(fromSquare, toSquare, piece.QueenType)
	move2 := NewMove(fromSquare, toSquare, piece.BishopType)

	if move1.Equals(move2) {
		t.Errorf("Expected %v to not equal %v", move1, move2)
	}
}
