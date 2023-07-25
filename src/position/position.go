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
func NewPosition(fen FEN) (*Position, error) {
	position, err := getPositionFromFEN(fen)
	return &position, err
}

func getPositionFromFEN(fen FEN) Position {
	white, black := getPiecePositionsFromFEN(fen)
	return Position{White: white, Black: black}
}

func getPiecePositionsFromFEN(fen FEN) (map[board.Square]pieces.Piece, map[board.Square]pieces.Piece) {
	white := make(map[board.Square]pieces.Piece)
	black := make(map[board.Square]pieces.Piece)

	for i, square := range board.Squares {
		piece, err := fen.GetPiece(square)
		if err != nil {
			continue
		}
		if piece.IsWhite() {
			white[i] = piece
		} else {
			black[i] = piece
		}
	}

	return white, black
}

// GetPiece returns the piece at the given square.
