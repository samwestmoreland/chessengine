package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Pawn struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewPawn(currentSquare board.Square, colour board.Colour) *Pawn {
	return &Pawn{CurrentSquare: currentSquare, Colour: colour}
}

// Returns the piece's color
func (p *Pawn) GetColour() board.Colour {
	return p.Colour
}

// Returns the piece's type
func (p *Pawn) Type() Type {
	panic("not implemented") // TODO: Implement
}

func (p *Pawn) GetLegalMoves() []board.Square {
	return []board.Square{}
}

func (p *Pawn) GetCurrentSquare() board.Square {
	return p.CurrentSquare
}
