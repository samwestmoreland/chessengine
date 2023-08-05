package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
)

// Queen is a piece that can move any number of squares diagonally, horizontally, or vertically
type Queen struct {
	CurrentSquare *board.Square
	Colour        board.Colour
}

// NewQueen creates a new queen piece
func NewQueen(square *board.Square, colour board.Colour) *Queen {
	return &Queen{
		CurrentSquare: square,
		Colour:        colour,
	}
}

// GetColour returns the piece's color
func (q *Queen) GetColour() board.Colour {
	return q.Colour
}

// Type returns the piece's type
func (q *Queen) Type() Type {
	return QueenType
}

// GetCurrentSquare returns the piece's current square
func (q *Queen) GetCurrentSquare() *board.Square {
	panic("not implemented") // TODO: Implement
}

// GetMoves returns a list of valid moves for the piece
func (q *Queen) GetMoves(board.Square, *Position) ([]moves.Move, error) {
	panic("not implemented") // TODO: Implement
}
