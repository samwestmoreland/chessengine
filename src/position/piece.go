package position

import (
	"unicode"

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
	GetMoves(*Position) ([]moves.Move, error)
}

// FromChar returns a piece from a character.
func FromChar(ch rune, square *board.Square) Piece {
	colour := getCase(ch)

	ch = unicode.ToLower(ch)
	switch ch {
	case 'k':
		return NewKing(square, colour)
	case 'q':
		return NewQueen(square, colour)
	case 'r':
		return NewRook(square, colour)
	case 'b':
		return NewBishop(square, colour)
	case 'n':
		return NewKnight(square, colour)
	case 'p':
		return NewPawn(square, colour)
	default:
		return nil
	}
}

func getCase(ch rune) board.Colour {
	if ch >= 'A' && ch <= 'Z' {
		return board.White
	} else if ch >= 'a' && ch <= 'z' {
		return board.Black
	}

	return board.Unknown
}
