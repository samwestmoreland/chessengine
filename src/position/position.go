package position

import (
	"fmt"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
)

// A Position represents a chess position.
type Position struct {
	White map[board.Square]Piece
	Black map[board.Square]Piece
	Turn  board.Colour
}

// NewPosition returns a new Position.
func NewPosition(fen *FEN) *Position {
	position := getPositionFromFEN(fen)
	return position
}

func getPositionFromFEN(fen *FEN) *Position {
	white, black := getPiecePositionsFromFEN(fen)
	ret := Position{White: white, Black: black}
	return &ret
}

func getPiecePositionsFromFEN(fen *FEN) (map[board.Square]Piece, map[board.Square]Piece) {
	white := make(map[board.Square]Piece)
	black := make(map[board.Square]Piece)

	for _, square := range board.Squares {
		fmt.Println("Calling GetPiece with square: ", square)
		piece, err := fen.GetPiece(square)
		if err != nil {
			continue
		}
		if piece == nil {
			continue
		}
		if piece.GetColour() == board.White {
			white[square] = piece
		} else {
			black[square] = piece
		}
	}

	return white, black
}

// GetAllPossibleMoves returns all possible moves for the current position.
func (p *Position) GetAllPossibleMoves() ([]moves.Move, error) {
	var moves []moves.Move
	if p.White == nil || p.Black == nil {
		return moves, fmt.Errorf("Position is not valid")
	}

	if p.Turn == board.White {
		return p.getWhiteMoves()
	}

	if p.Turn == board.Black {
		return p.getBlackMoves()
	}

	return moves, fmt.Errorf("Position is not valid")
}

func (p *Position) getWhiteMoves() ([]moves.Move, error) {
	var moves []moves.Move
	for square, piece := range p.White {
		moves = append(moves, piece.GetMoves(square, p))
	}
	return moves, nil
}

func (p *Position) getBlackMoves() ([]moves.Move, error) {
	var moves []moves.Move
	for square, piece := range p.Black {
		moves = append(moves, piece.GetMoves(square, p))
	}
	return moves, nil
}

// getPiece returns the piece at the given square, or an error if the square is invalid.
func (p *Position) getPiece(square board.Square) (Piece, error) {
	if err := square.Valid(); err != nil {
		return nil, err
	}
	if piece, ok := p.White[square]; ok {
		return piece, nil
	}
	if piece, ok := p.Black[square]; ok {
		return piece, nil
	}
	return nil, nil
}
