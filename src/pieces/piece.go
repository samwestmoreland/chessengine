package pieces

import (
	"fmt"
)

type Piece interface {
	// Returns the piece's color
	Color() Color
	// Returns the piece's type
	Type() Type
	// Returns the piece's current square
	Square() Square
	LegalMoves() []Move
}
