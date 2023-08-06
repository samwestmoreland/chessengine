package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Knight represents a knight piece.
type Knight struct {
	CurrentSquare *board.Square
	Colour        board.Colour
}

// NewKnight creates a new knight.
func NewKnight(currentSquare *board.Square, colour board.Colour) *Knight {
	return &Knight{CurrentSquare: currentSquare, Colour: colour}
}

// GetColour returns the piece's color.
func (k *Knight) GetColour() board.Colour {
	return k.Colour
}

// Type returns the piece's type.
func (k *Knight) Type() piece.Type {
	return piece.KnightType
}

// GetCurrentSquare returns the piece's current square.
func (k *Knight) GetCurrentSquare() *board.Square {
	panic("not implemented") // TODO: Implement
}

// GetMoves returns a list of valid moves for the piece.
func (k *Knight) GetMoves(board.Square, *Position) ([]moves.Move, error) {
	panic("not implemented") // TODO: Implement
}
