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
func (b *Bishop) GetMoves(pos *Position) (moves.MoveList, error) {
	ret := moves.MoveList{}
	// there are 4 directions a bishop can move in
	// we'll iterate over each direction and add the valid moves
	// until we hit a piece or the edge of the board
	for _, direction := range []board.Direction{board.NorthEast, board.NorthWest, board.SouthEast, board.SouthWest} {
		oldSquare := b.CurrentSquare
		for i := 1; i < 7; i++ {
			newSquare := oldSquare.Translate(direction)
			if !newSquare.Valid() {
				break
			}

			squareIsOccupied, col := pos.squareIsOccupied(newSquare)
			if squareIsOccupied && col == b.GetColour() {
				break
			} else if squareIsOccupied && col != b.GetColour() {
				break
			}

			ret.AddMove(moves.NewMove(b.CurrentSquare, newSquare, piece.BishopType, false))
			oldSquare = newSquare
		}
	}

	return ret, nil
}
