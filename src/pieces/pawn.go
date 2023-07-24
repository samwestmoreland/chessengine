package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Pawn struct {
	CurrentSquare board.Square
}

func (p *Pawn) GetLegalMoves() []board.Square {
	return []board.Square{}
}

func (p *Pawn) GetCurrentSquare() board.Square {
	return p.CurrentSquare
}
