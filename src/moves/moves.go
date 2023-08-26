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
	Capture   bool
}

// NewMove creates a new Move.
func NewMove(from board.Square, to board.Square, pieceType piece.Type, capture bool) Move {
	return Move{From: from, To: to, PieceType: pieceType, Capture: capture}
}

func (m Move) String() string {
	ret := m.PieceType.String() + ": " + m.From.String() + " -> " + m.To.String()
	if m.Capture {
		ret += " (capture)"
	}
	return ret
}

// Equals returns true if two moves are equal.
func (m Move) Equals(other Move) bool {
	return m.From == other.From &&
		m.To == other.To &&
		m.PieceType == other.PieceType &&
		m.Capture == other.Capture
}

func (m Move) IsCapture() bool {
	return m.Capture
}

func (m Move) GetPieceType() piece.Type {
	return m.PieceType
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
