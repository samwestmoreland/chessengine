package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

// Pawn is a piece that can move one square forward, or two squares forward
// if it is on its starting square, and can capture diagonally.
type Pawn struct {
	CurrentSquare *board.Square
	Colour        board.Colour
}

// NewPawn creates a new pawn
func NewPawn(currentSquare *board.Square, colour board.Colour) *Pawn {
	return &Pawn{CurrentSquare: currentSquare, Colour: colour}
}

// GetColour returns the piece's color
func (p *Pawn) GetColour() board.Colour {
	return p.Colour
}

// Type returns the piece's type
func (p *Pawn) Type() piece.Type {
	return piece.PawnType
}

// GetCurrentSquare returns the piece's current square
func (p *Pawn) GetCurrentSquare() *board.Square {
	return p.CurrentSquare
}

// GetMoves returns a list of moves that the piece can make
func (p *Pawn) GetMoves(board.Square, *Position) ([]moves.Move, error) {
	ret := make([]moves.Move, 0, 4)

	return ret, nil
}
