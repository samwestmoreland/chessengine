package position

import (
	"github.com/samwestmoreland/chessengine/internal/board"
	"github.com/samwestmoreland/chessengine/internal/moves"
	"github.com/samwestmoreland/chessengine/internal/piece"
)

// King is a struct representing a king piece.
type King struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

// NewKing returns a new king piece.
func NewKing(currentSquare board.Square, colour board.Colour) *King {
	return &King{CurrentSquare: currentSquare, Colour: colour}
}

// GetColour returns the piece's color.
func (k *King) GetColour() board.Colour {
	return k.Colour
}

// Type returns the piece's type.
func (k *King) Type() piece.Type {
	return piece.KingType
}

// GetCurrentSquare returns the piece's current square.
func (k *King) GetCurrentSquare() board.Square {
	return k.CurrentSquare
}

// GetMoves returns a list of all possible moves for the king.
func (k *King) GetMoves(position *Position) (moves.MoveList, error) {
	// The king can move one square in any direction, so there are 8 possible moves
	ret := moves.MoveList{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			}

			// Get the square
			square := board.Square{File: k.CurrentSquare.File + i - 1, Rank: k.CurrentSquare.Rank + j - 1}
			if !square.Valid() {
				continue
			}

			// Check if the square is occupied by a friendly piece
			pieceAtSquare, err := position.getPiece(square)
			if err != nil {
				return moves.MoveList{}, err
			}

			if pieceAtSquare == nil {
				m := moves.NewMove(k.CurrentSquare, square, piece.KingType, false)
				ret.AddMove(m)

				continue
			}

			if (pieceAtSquare).GetColour() == k.Colour {
				continue
			}

			// The square is not occupied by a friendly piece, add it to the list
			m := moves.NewMove(k.CurrentSquare, square, piece.KingType, false)
			ret.AddMove(m)
		}
	}

	return ret, nil
}
