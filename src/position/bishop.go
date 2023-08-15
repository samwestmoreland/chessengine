// Package position provides a representation of a chess position, as well as the logic for the various chess pieces
package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Bishop represents a bishop piece.
type Bishop struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

// NewBishop creates a new bishop piece.
func NewBishop(square board.Square, colour board.Colour) *Bishop {
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
func (b *Bishop) GetCurrentSquare() board.Square {
	return b.CurrentSquare
}

// GetMoves returns the piece's valid moves.
func (b *Bishop) GetMoves(*Position) ([]moves.Move, error) {
	ret := make([]moves.Move, 0, 14)
	// there are 4 directions a bishop can move in
	// we'll iterate over each direction and add the valid moves
	// until we hit a piece or the edge of the board
	for _, direction := range []board.Direction{board.NorthEast, board.NorthWest, board.SouthEast, board.SouthWest} {
		for square := b.CurrentSquare.Translate(direction); square.Valid(); square = square.Translate(direction) {
			ret = append(ret, moves.NewMove(b.CurrentSquare, square, piece.BishopType))
		}
	}

	return ret, nil
}
