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
