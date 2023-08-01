package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
)

type Queen struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewQueen(square board.Square, colour board.Colour) *Queen {
	return &Queen{
		CurrentSquare: square,
		Colour:        colour,
	}
}

// Returns the piece's color
func (q *Queen) GetColour() board.Colour {
	return q.Colour
}

// Returns the piece's type
func (q *Queen) Type() Type {
	return QueenType
}

// Returns the piece's current square
func (q *Queen) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (q *Queen) GetMoves(board.Square, *Position) []moves.Move {
	panic("not implemented") // TODO: Implement
}
