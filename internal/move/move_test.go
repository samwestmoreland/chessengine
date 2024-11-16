package move_test

import (
	"testing"

	"github.com/samwestmoreland/chessengine/internal/move"
	"github.com/samwestmoreland/chessengine/internal/piece"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

func TestExtractSourceAndTargetSquares(t *testing.T) {
	aMove := move.Encode(sq.A1, sq.A2, piece.Wk, 0, 0, 0, 0, 0)

	if aMove.Source() != sq.A1 {
		t.Errorf("Expected source square to be a1, got %s", sq.Stringify(aMove.Source()))
	}

	if aMove.Target() != sq.A2 {
		t.Errorf("Expected target square to be a2, got %s", sq.Stringify(aMove.Target()))
	}
}
