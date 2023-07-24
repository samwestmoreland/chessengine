package pieces

import (
	"fmt"

	"github.com/samwestmoreland/chessengine/board"
)

type Piece interface {
	// Returns the piece's color
	Color() board.Color
	// Returns the piece's type
	Type() Type
	// Returns the piece's current square
	Square() Square
	LegalMoves() []Move
}
