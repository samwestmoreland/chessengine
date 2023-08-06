package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Piece represents a chess piece.
type Piece interface {
	// Returns the piece's color
	GetColour() board.Colour
	// Returns the piece's type
	Type() piece.Type
	// Returns the piece's current square
	GetCurrentSquare() *board.Square
	GetMoves(board.Square, *Position) ([]moves.Move, error)
}

// FromChar returns a piece from a character.
func FromChar(ch rune, square *board.Square) Piece {
	switch ch {
	case 'K':
		return NewKing(square, board.White)
	case 'k':
		return NewKing(square, board.Black)
	case 'Q':
		return NewQueen(square, board.White)
	case 'q':
		return NewQueen(square, board.Black)
	case 'R':
		return NewRook(square, board.White)
	case 'r':
		return NewRook(square, board.Black)
	case 'B':
		return NewBishop(square, board.White)
	case 'b':
		return NewBishop(square, board.Black)
	case 'N':
		return NewKnight(square, board.White)
	case 'n':
		return NewKnight(square, board.Black)
	case 'P':
		return NewPawn(square, board.White)
	case 'p':
		return NewPawn(square, board.Black)
	default:
		return nil
	}
}
