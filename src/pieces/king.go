package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type King struct {
	CurrentSqure board.Square
}

func NewKing(currentSquare board.Square) *King {
	return &King{CurrentSqure: currentSquare}
}

// Returns the piece's color
func (k *King) Colour() board.Colour {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's type
func (k *King) Type() Type {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's current square
func (k *King) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (k *King) GetLegalMoves() []board.Square {
	panic("not implemented") // TODO: Implement
}
