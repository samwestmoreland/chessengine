package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
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
	return PawnType
}

func (p *Pawn) GetCurrentSquare() board.Square {
	return p.CurrentSquare
}

func (p *Pawn) GetMoves(board.Square, *Position) []moves.Move {
	panic("implement me")
}
