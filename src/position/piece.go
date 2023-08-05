package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
)

// Type represents the type of a piece
type Type int

const (
	KingType   Type = iota // KingType = 0
	QueenType              // QueenType = 1
	RookType               // RookType = 2
	BishopType             // BishopType = 3
	KnightType             // KnightType = 4
	PawnType               // PawnType = 5
	NoneType               // NoneType = 6
)

// Piece represents a chess piece
type Piece interface {
	// Returns the piece's color
	GetColour() board.Colour
	// Returns the piece's type
	Type() Type
	// Returns the piece's current square
	GetCurrentSquare() *board.Square
	GetMoves(board.Square, *Position) ([]moves.Move, error)
}

// FromChar returns a piece from a character
func FromChar(ch rune, sq *board.Square) Piece {
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
