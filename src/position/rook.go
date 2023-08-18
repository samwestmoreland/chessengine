package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
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
func (r *Rook) GetMoves(pos *Position) ([]moves.Move, error) {
	ret := []moves.Move{}

	for _, direction := range []board.Direction{board.North, board.East, board.South, board.West} {
		oldSquare := r.CurrentSquare
		for i := 1; i < 8; i++ {
			newSquare := oldSquare.Translate(direction)
			if !newSquare.Valid() {
				break
			}
			pieceOnSquare, err := pos.getPiece(newSquare)
			if err != nil {
				return nil, err
			}

			if pieceOnSquare != nil {
				if pieceOnSquare.GetColour() == r.GetColour() {
					break
				}
			}

			ret = append(ret, moves.NewMove(r.CurrentSquare, newSquare, piece.RookType))
			oldSquare = newSquare

		}
	}

	return ret, nil
}
