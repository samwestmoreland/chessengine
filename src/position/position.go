package position

import (
	"fmt"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/pieces"
)

// A Position represents a chess position.
type Position struct {
	White map[board.Square]pieces.Piece
	Black map[board.Square]pieces.Piece
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

func getPiecePositionsFromFEN(fen *FEN) (map[board.Square]pieces.Piece, map[board.Square]pieces.Piece) {
	white := make(map[board.Square]pieces.Piece)
	black := make(map[board.Square]pieces.Piece)

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
