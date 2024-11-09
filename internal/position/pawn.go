package position

import (
	"github.com/samwestmoreland/chessengine/internal/board"
	"github.com/samwestmoreland/chessengine/internal/moves"
	"github.com/samwestmoreland/chessengine/internal/piece"
)

// Pawn is a piece that can move one square forward, or two squares forward
// if it is on its starting square, and can capture diagonally.
type Pawn struct {
	CurrentSquare board.Square
	Colour        board.Colour
}

// NewPawn creates a new pawn.
func NewPawn(currentSquare board.Square, colour board.Colour) *Pawn {
	return &Pawn{CurrentSquare: currentSquare, Colour: colour}
}

// GetColour returns the piece's color.
func (p Pawn) GetColour() board.Colour {
	return p.Colour
}

// Type returns the piece's type.
func (p Pawn) Type() piece.Type {
	return piece.PawnType
}

// GetCurrentSquare returns the piece's current square.
func (p Pawn) GetCurrentSquare() board.Square {
	return p.CurrentSquare
}

// GetMoves returns a list of moves that the piece can make.
func (p Pawn) GetMoves(pos *Position) (moves.MoveList, error) {
	ret := moves.MoveList{}

	if p.Colour == board.White {
		ret.AddMoves(p.getForwardMovesWhite(pos))
	} else if p.Colour == board.Black {
		ret.AddMoves(p.getForwardMovesBlack(pos))
	}

	return ret, nil
}

func (p *Pawn) getForwardMovesWhite(pos *Position) []moves.Move {
	ret := make([]moves.Move, 0, 2)

	// Move one square forward.
	destination := p.CurrentSquare.Translate(board.North)
	if occ, _ := pos.squareIsOccupied(destination); destination.Valid() && !occ {
		ret = append(ret, moves.NewMove(p.CurrentSquare, destination, piece.PawnType, false))
	}

	// Move two squares forward.
	if p.CurrentSquare.Rank == 2 {
		destination = destination.Translate(board.North)
		if occ, _ := pos.squareIsOccupied(destination); destination.Valid() && !occ {
			ret = append(ret, moves.NewMove(p.CurrentSquare, destination, piece.PawnType, false))
		}
	}

	return ret
}

func (p *Pawn) getForwardMovesBlack(pos *Position) []moves.Move {
	ret := make([]moves.Move, 0, 2)

	// Move one square forward.
	destination := p.CurrentSquare.Translate(board.South)
	if occ, _ := pos.squareIsOccupied(destination); destination.Valid() && !occ {
		ret = append(ret, moves.NewMove(p.CurrentSquare, destination, piece.PawnType, false))
	}

	// Move two squares forward.
	if p.CurrentSquare.Rank == 7 {
		destination = destination.Translate(board.South)
		if occ, _ := pos.squareIsOccupied(destination); destination.Valid() && !occ {
			ret = append(ret, moves.NewMove(p.CurrentSquare, destination, piece.PawnType, false))
		}
	}

	return ret
}

func (p *Pawn) IsDoubled(pos *Position) bool {
	piecesOfSameColour := map[board.Square]Piece{}

	if p.Colour == board.White {
		piecesOfSameColour = pos.GetWhitePieces()
	} else if p.Colour == board.Black {
		piecesOfSameColour = pos.GetBlackPieces()
	}

	for square, pi := range piecesOfSameColour {
		if pi.Type() == piece.PawnType && square.File == p.CurrentSquare.File {
			return true
		}
	}

	return false
}

func (p *Pawn) IsIsolated(pos *Position) bool {
	piecesOfSameColour := map[board.Square]Piece{}

	if p.Colour == board.White {
		piecesOfSameColour = pos.GetWhitePieces()
	} else if p.Colour == board.Black {
		piecesOfSameColour = pos.GetBlackPieces()
	}

	for square, pi := range piecesOfSameColour {
		if pi.Type() == piece.PawnType && square.File == p.CurrentSquare.File-1 {
			return false
		} else if pi.Type() == piece.PawnType && square.File == p.CurrentSquare.File+1 {
			return false
		}
	}

	return true
}
