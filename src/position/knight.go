package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Knight represents a knight piece.
type Knight struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

// NewKnight creates a new knight.
func NewKnight(currentSquare board.Square, colour board.Colour) *Knight {
	return &Knight{CurrentSquare: currentSquare, Colour: colour}
}

// GetColour returns the piece's color.
func (k *Knight) GetColour() board.Colour {
	return k.Colour
}

// Type returns the piece's type.
func (k *Knight) Type() piece.Type {
	return piece.KnightType
}

// GetCurrentSquare returns the piece's current square.
func (k *Knight) GetCurrentSquare() board.Square {
	return k.CurrentSquare
}

// GetMoves returns a list of valid moves for the piece.
func (k *Knight) GetMoves(pos *Position) ([]moves.Move, error) {
	ret := make([]moves.Move, 0, 8)

	// we can iterate over 2 and -2 and 1 and -1 to get all the possible moves
	for _, xOffset := range []int{2, -2} {
		for _, yOffset := range []int{1, -1} {
			newSquare := board.Square{Rank: k.CurrentSquare.Rank + xOffset, File: k.CurrentSquare.File + yOffset}
			if !newSquare.Valid() {
				continue
			}

			if occ, col := pos.squareIsOccupied(newSquare); !occ || col != k.Colour {
				m := moves.NewMove(k.CurrentSquare, newSquare, piece.KnightType)
				ret = append(ret, m)
			}

			newSquare = board.Square{Rank: k.CurrentSquare.Rank + yOffset, File: k.CurrentSquare.File + xOffset}
			if !newSquare.Valid() {
				continue
			}

			if occ, col := pos.squareIsOccupied(newSquare); !occ || col != k.Colour {
				m := moves.NewMove(k.CurrentSquare, newSquare, piece.KnightType)
				ret = append(ret, m)
			}
		}
	}

	return ret, nil
}
