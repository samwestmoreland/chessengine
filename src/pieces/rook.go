package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Rook struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewRook(colour board.Colour, square board.Square) *Rook {
	return &Rook{
		CurrentSquare: square,
		Colour:        colour,
	}
}

// Returns the piece's color
func (q *Rook) GetColour() board.Colour {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's type
func (q *Rook) Type() Type {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's current square
func (q *Rook) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (q *Rook) GetLegalMoves() []board.Square {
	panic("not implemented") // TODO: Implement
}
