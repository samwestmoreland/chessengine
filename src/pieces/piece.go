package pieces

import (
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
	NoneType
)

type Piece interface {
	// Returns the piece's color
	Colour() board.Colour
	// Returns the piece's type
	Type() Type
	// Returns the piece's current square
	GetCurrentSquare() board.Square
	GetLegalMoves() []board.Square
}
