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

func TestClearlyWinningPositionForBlack(t *testing.T) {
	fen, err := position.ParseFEN("5k2/7p/1p4pb/8/8/1bN4P/1P4KP/8 b - - 1 28")
	if err != nil {
		t.Error(err)
	}

	pos := position.NewPositionFromFEN(fen)
	evaluator := NewShannonEvaluator()

	if score := evaluator.Evaluate(pos); score > 0 {
		t.Errorf("Expected score to be negative, got %v", score)
	}
}
