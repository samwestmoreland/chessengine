package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
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
	GetColour() board.Colour
	// Returns the piece's type
	Type() Type
	// Returns the piece's current square
	GetCurrentSquare() board.Square
	GetMoves(board.Square, *Position) []moves.Move
}

func FromChar(ch rune, sq board.Square) Piece {
	switch ch {
	case 'K':
		return NewKing(sq, board.White)
	case 'k':
		return NewKing(sq, board.Black)
	case 'Q':
		return NewQueen(sq, board.White)
	case 'q':
		return NewQueen(sq, board.Black)
	case 'R':
		return NewRook(sq, board.White)
	case 'r':
		return NewRook(sq, board.Black)
	case 'B':
		return NewBishop(sq, board.White)
	case 'b':
		return NewBishop(sq, board.Black)
	case 'N':
		return NewKnight(sq, board.White)
	case 'n':
		return NewKnight(sq, board.Black)
	case 'P':
		return NewPawn(sq, board.White)
	case 'p':
		return NewPawn(sq, board.Black)
	default:
		return nil
	}
}
