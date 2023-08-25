package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Queen is a piece that can move any number of squares diagonally, horizontally, or vertically.
type Queen struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

// NewQueen creates a new queen piece.
func NewQueen(square board.Square, colour board.Colour) *Queen {
	return &Queen{
		CurrentSquare: square,
		Colour:        colour,
	}
}

// GetColour returns the piece's color.
func (q *Queen) GetColour() board.Colour {
	return q.Colour
}

// Type returns the piece's type.
func (q *Queen) Type() piece.Type {
	return piece.QueenType
}

// GetCurrentSquare returns the piece's current square.
func (q *Queen) GetCurrentSquare() board.Square {
	return q.CurrentSquare
}

// GetMoves returns a list of valid moves for the piece.
func (q *Queen) GetMoves(pos *Position) ([]moves.Move, error) {
	ret := make([]moves.Move, 0, 21)

	for _, direction := range []board.Direction{
		board.North,
		board.NorthEast,
		board.East,
		board.SouthEast,
		board.South,
		board.SouthWest,
		board.West,
		board.NorthWest,
	} {
		oldSquare := q.CurrentSquare
		for i := 0; i < 7; i++ {
			newSquare := oldSquare.Translate(direction)
			if !newSquare.Valid() {
				break
			}

			squareIsOccupied, col := pos.squareIsOccupied(newSquare)
			if squareIsOccupied && col == q.GetColour() {
				break
			} else if squareIsOccupied && col != q.GetColour() {
				ret = append(ret, moves.NewMove(q.CurrentSquare, newSquare, piece.QueenType, true))
				break
			}

			ret = append(ret, moves.NewMove(q.CurrentSquare, newSquare, piece.QueenType, false))
			oldSquare = newSquare
		}
	}

	return ret, nil
}
