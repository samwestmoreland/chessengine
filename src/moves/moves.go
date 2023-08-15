package moves

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Move represents a move from one square to another.
type Move struct {
	From      board.Square
	To        board.Square
	PieceType piece.Type
}

// NewMove creates a new Move.
func NewMove(from board.Square, to board.Square, pieceType piece.Type) Move {
	return Move{From: from, To: to, PieceType: pieceType}
}

func (m Move) String() string {
	return m.PieceType.String() + ": " + m.From.String() + " -> " + m.To.String()
}

// Equals returns true if two moves are equal.
func (m Move) Equals(other Move) bool {
	return m.From == other.From && m.To == other.To && m.PieceType == other.PieceType
}

func MoveListsEqual(moves1 []Move, moves2 []Move) bool {
	// Check that the moves are the same, but don't care about order
	for _, move1 := range moves1 {
		found := false

		for _, move2 := range moves2 {
			if move1.Equals(move2) {
				found = true

				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
