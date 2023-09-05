package eval

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/position"
)

func TestClearlyWinningPositionForWhite(t *testing.T) {
	fen, err := position.ParseFEN("7b/1p3p2/4pkp1/1BB4p/5P2/8/PP3P1P/6K1 b - - 0 24")
	if err != nil {
		t.Error(err)
	}

	pos := position.NewPositionFromFEN(fen)
	evaluator := NewShannonEvaluator()

	if score := evaluator.Evaluate(pos); score < 0 {
		t.Errorf("Expected score to be positive, got %v", score)
	}
}
