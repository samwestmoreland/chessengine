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

func FromChar(char rune) Piece {
	switch char {
	case 'K':
		return NewKing(board.White)
	case 'k':
		return NewKing(board.Black)
	case 'Q':
		return NewQueen(board.White)
	case 'q':
		return NewQueen(board.Black)
	case 'R':
		return NewRook(board.White)
	case 'r':
		return NewRook(board.Black)
	case 'B':
		return NewBishop(board.White)
	case 'b':
		return NewBishop(board.Black)
	case 'N':
		return NewKnight(board.White)
	case 'n':
		return NewKnight(board.Black)
	case 'P':
		return NewPawn(board.White)
	case 'p':
		return NewPawn(board.Black)
	default:
		return nil
	}
}
