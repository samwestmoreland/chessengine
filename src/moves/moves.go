package moves

import (
	"github.com/samwestmoreland/chessengine/src/board"
)

type Move struct {
	From      board.Square
	To        board.Square
	PieceType string
}

func (m Move) String() string {
	return m.From.String() + " -> " + m.To.String()
}
