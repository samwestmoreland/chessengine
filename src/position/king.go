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

func (k *King) GetMoves(board.Square, *Position) []moves.Move {
	panic("not implemented") // TODO: Implement
}
