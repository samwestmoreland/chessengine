package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
)

// Rook represents a rook piece, which can move horizontally or vertically
type Rook struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

// NewRook creates a new rook piece
func NewRook(square board.Square, colour board.Colour) *Rook {
	return &Rook{
		CurrentSquare: square,
		Colour:        colour,
	}
}

// GetColour returns the piece's color
func (r *Rook) GetColour() board.Colour {
	return r.Colour
}

// Type returns the piece's type
func (r *Rook) Type() Type {
	return RookType
}

// GetCurrentSquare returns the piece's current square
func (r *Rook) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

// GetMoves returns the piece's valid moves
func (r *Rook) GetMoves(board.Square, *Position) ([]moves.Move, error) {
	panic("not implemented") // TODO: Implement
}
