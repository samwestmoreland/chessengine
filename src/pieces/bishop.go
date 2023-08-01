package pieces

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/position"
)

type Bishop struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

func NewBishop(square board.Square, colour board.Colour) *Bishop {
	return &Bishop{CurrentSquare: square, Colour: colour}
}

// Returns the piece's color
func (b *Bishop) GetColour() board.Colour {
	return b.Colour
}

// Returns the piece's type
func (b *Bishop) Type() Type {
	return BishopType
}

// Returns the piece's current square
func (b *Bishop) GetCurrentSquare() board.Square {
	panic("not implemented") // TODO: Implement
}

func (b *Bishop) GetMoves(board.Square, *position.Position) []moves.Move {
	panic("not implemented") // TODO: Implement
}
