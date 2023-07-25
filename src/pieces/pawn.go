package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Pawn struct {
	CurrentSquare board.Square
}

func NewPawn(currentSquare board.Square) *Pawn {
	return &Pawn{CurrentSquare: currentSquare}
}

// Returns the piece's color
func (p *Pawn) Colour() board.Colour {
	panic("not implemented") // TODO: Implement
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
