package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// NoPiece is a struct representing an empty square.
type NoPiece struct {
	CurrentSquare *board.Square
	Colour        board.Colour
}

// NewNoPiece returns a new king piece.
func NewNoPiece(currentSquare *board.Square, colour board.Colour) *NoPiece {
	return &NoPiece{CurrentSquare: currentSquare, Colour: colour}
}

// GetColour returns the piece's color.
func (k *NoPiece) GetColour() board.Colour {
	return k.Colour
}

// Type returns the piece's type.
func (k *NoPiece) Type() piece.Type {
	return piece.NoneType
}

// GetCurrentSquare returns the piece's current square.
func (k *NoPiece) GetCurrentSquare() *board.Square {
	return k.CurrentSquare
}

// GetMoves returns a list of all possible moves for the king.
func (k *NoPiece) GetMoves(_ *Position) ([]moves.Move, error) {
	return []moves.Move{}, nil
}
