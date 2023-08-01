package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type King struct {
	CurrentSqure board.Square
	Colour       board.Colour
}

func NewKing(currentSquare board.Square, colour board.Colour) *King {
	return &King{CurrentSqure: currentSquare, Colour: colour}
}

// Returns the piece's color
func (k *King) GetColour() board.Colour {
	return k.Colour
}

// Returns the piece's type
func (k *King) Type() Type {
	return KingType
}

// Returns the piece's current square
func (k *King) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (k *King) GetLegalMoves() []board.Square {
	panic("not implemented") // TODO: Implement
}
