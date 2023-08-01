package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/position"
)

type Rook struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewRook(square board.Square, colour board.Colour) *Rook {
	return &Rook{
		CurrentSquare: square,
		Colour:        colour,
	}
}

// Returns the piece's color
func (r *Rook) GetColour() board.Colour {
	return r.Colour
}

// Returns the piece's type
func (r *Rook) Type() Type {
	return RookType
}

// Returns the piece's current square
func (r *Rook) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (r *Rook) GetMoves(board.Square, *position.Position) []moves.Move {
	panic("not implemented") // TODO: Implement
}
