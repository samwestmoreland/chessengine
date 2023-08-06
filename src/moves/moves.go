package moves

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/piece"
)

type Move struct {
	From      *board.Square
	To        *board.Square
	PieceType piece.Type
}

func (m Move) String() string {
	return m.PieceType.String() + ": " + m.From.String() + " -> " + m.To.String()
}
}
