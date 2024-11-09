package position

import (
	"github.com/samwestmoreland/chessengine/internal/board"
	"github.com/samwestmoreland/chessengine/internal/moves"
	"github.com/samwestmoreland/chessengine/internal/piece"
)

// Rook represents a rook piece, which can move horizontally or vertically.
type Rook struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

// NewRook creates a new rook piece.
func NewRook(square board.Square, colour board.Colour) *Rook {
	return &Rook{
		CurrentSquare: square,
		Colour:        colour,
	}
}

// GetColour returns the piece's colour.
func (r *Rook) GetColour() board.Colour {
	return r.Colour
}

// Type returns the piece's type.
func (r *Rook) Type() piece.Type {
	return piece.RookType
}

// GetCurrentSquare returns the piece's current square.
func (r *Rook) GetCurrentSquare() board.Square {
	return r.CurrentSquare
}

// GetMoves returns the piece's valid moves.
func (r *Rook) GetMoves(pos *Position) (moves.MoveList, error) {
	ret := moves.MoveList{}

	for _, direction := range []board.Direction{board.North, board.East, board.South, board.West} {
		oldSquare := r.CurrentSquare
		for i := 1; i < 7; i++ {
			newSquare := oldSquare.Translate(direction)
			if !newSquare.Valid() {
				break
			}

			squareIsOccupied, col := pos.squareIsOccupied(newSquare)
			if squareIsOccupied && col == r.GetColour() {
				break
			} else if squareIsOccupied && col != r.GetColour() {
				ret.AddMove(moves.NewMove(r.CurrentSquare, newSquare, piece.RookType, true))

				break
			}

			ret.AddMove(moves.NewMove(r.CurrentSquare, newSquare, piece.RookType, false))
			oldSquare = newSquare
		}
	}

	return ret, nil
}
