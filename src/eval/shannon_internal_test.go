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

func TestMaterialScore(t *testing.T) {
	expectedResults := map[string]int{
		"r2q1rk1/ppp1bppp/3pp1n1/1B6/1n1PP2P/2N2PB1/P1P1QP2/R3K2R b KQ - 1 13": -1,
		"r4rk1/4b1pp/n2pp1p1/qpp5/p2PP3/P1N1QPB1/B1P2P2/R3K2R w KQ - 0 21":     1,
		"r7/6pk/n3B1pp/q3Pr2/p3N2R/p3QP2/2P1KP2/7R w - - 1 30":                 2,
		"8/8/2BPnk2/b1P3p1/2K2pp1/6P1/5B1P/4q3 w - - 0 50":                     -8,
		"8/8/1p6/6K1/7p/1P4k1/8/6b1 w - - 2 41":                                -4,
	}

	ev := NewShannonEvaluator()

	for fen, expectedScore := range expectedResults {
		parsed, _ := position.ParseFEN(fen)
		pos := position.NewPositionFromFEN(parsed)
		if score := ev.getMaterialScore(pos); score != expectedScore {
			t.Errorf("Expected material score to be %v, got %v", expectedScore, score)
		}
	}

}
