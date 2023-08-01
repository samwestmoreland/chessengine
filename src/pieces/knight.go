package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Knight struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewKnight(currentSquare board.Square, colour board.Colour) *Knight {
	return &Knight{CurrentSquare: currentSquare}
}

// Returns the piece's color
func (k *Knight) GetColour() board.Colour {
	return k.Colour
}

// Returns the piece's type
func (k *Knight) Type() Type {
	return KnightType
}

// Returns the piece's current square
func (k *Knight) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (k *Knight) GetLegalMoves() []board.Square {
	panic("not implemented") // TODO: Implement
}
