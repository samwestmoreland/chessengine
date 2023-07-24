package pieces

import (
	"fmt"

	"github.com/samwestmoreland/chessengine/src/board"
)

type Type int

const (
	KingType Type = iota
	QueenType
	RookType
	BishopType
	KnightType
	PawnType
)

type Move struct {
	From board.Square
	To   board.Square
}

type Piece interface {
	// Returns the piece's color
	Color() board.Color
	// Returns the piece's type
	Type() Type
	// Returns the piece's current square
	Square() board.Square
	LegalMoves() []Move
}
