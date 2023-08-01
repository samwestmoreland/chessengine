package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
)

type King struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewKing(currentSquare board.Square, colour board.Colour) *King {
	return &King{CurrentSquare: currentSquare, Colour: colour}
}

// Returns the piece's color
func (k *King) GetColour() board.Colour {
	return k.Colour
}

// Returns the piece's type
func (k *King) Type() Type {
	return KingType
}

// Returns the piece's current square
func (k *King) GetCurrentSquare() board.Square {
	return k.CurrentSquare
}

func (k *King) GetMoves(sq board.Square, p *Position) ([]moves.Move, error) {
	// Get all possible moves assuming no other pieces on the board
	// The king can move one square in any direction, so there are 8 possible moves
	ret := make([]moves.Move, 0, 8)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Skip the current square
			if i == 1 && j == 1 {
				continue
			}

			// Get the square
			s := board.Square{File: k.CurrentSquare.File + i - 1, Rank: k.CurrentSquare.Rank + j - 1}
			if err := s.Valid(); err != nil {
				continue
			}

			// Check if the square is occupied by a friendly piece
			piece, err := p.getPiece(sq)
			if err != nil {
				return nil, err
			}

			if piece.GetColour() == k.Colour {
				continue
			}

			// The square is not occupied by a friendly piece, add it to the list
			ret = append(ret, moves.Move{From: k.CurrentSquare, To: s})
		}
	}
	return ret, nil
}
