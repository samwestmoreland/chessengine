package position

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/pieces"
)

// A Position represents a chess position.
type Position struct {
	White map[board.Square]pieces.Piece
	Black map[board.Square]pieces.Piece
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
		piece, err := fen.GetPiece(square)
		if err != nil {
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
