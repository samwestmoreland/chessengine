package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Queen struct {
	CurrentSquare board.Square
}

// Returns the piece's color
func (q *Queen) Colour() board.Colour {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's type
func (q *Queen) Type() Type {
	panic("not implemented") // TODO: Implement
}

// Returns the piece's current square
func (q *Queen) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (q *Queen) GetLegalMoves() []board.Square {
	panic("not implemented") // TODO: Implement
}
