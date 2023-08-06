// Package position provides a representation of a chess position, as well as the logic for the various chess pieces
package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Bishop represents a bishop piece.
type Bishop struct {
	CurrentSquare *board.Square
	Colour        board.Colour
}

// NewBishop creates a new bishop piece.
func NewBishop(square *board.Square, colour board.Colour) *Bishop {
	return &Bishop{CurrentSquare: square, Colour: colour}
}

// GetColour returns the piece's color.
func (b *Bishop) GetColour() board.Colour {
	return b.Colour
}

// Type returns the piece's type.
func (b *Bishop) Type() piece.Type {
	return piece.BishopType
}

// GetCurrentSquare returns the piece's current square.
func (b *Bishop) GetCurrentSquare() *board.Square {
	panic("not implemented") // TODO: Implement
}

// GetMoves returns the piece's valid moves.
func (b *Bishop) GetMoves(board.Square, *Position) ([]moves.Move, error) {
	panic("not implemented") // TODO: Implement
}
