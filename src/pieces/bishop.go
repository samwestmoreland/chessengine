package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Bishop struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewBishop(square board.Square, colour board.Colour) *Bishop {
	return &Bishop{CurrentSquare: square, Colour: colour}
}

// Returns the piece's color
func (b *Bishop) GetColour() board.Colour {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's type
func (b *Bishop) Type() Type {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's current square
func (b *Bishop) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (b *Bishop) GetLegalMoves() []board.Square {
	panic("not implemented") // TODO: Implement
}
