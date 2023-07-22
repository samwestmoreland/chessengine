package position

import (
	"testing"
)

func TestValidPositionCheck(t *testing.T) {
	t.Log("Testing valid position check")
	var fen FEN = FEN{Raw: "r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3"}
	if !fen.IsValid() {
		t.Error("Expected valid position")
	}
}

func TestInvalidPosition(t *testing.T) {
	t.Log("Testing invalid position check")
	var fen FEN = FEN{Raw: "r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R"}
	if fen.IsValid() {
		t.Error("Expected invalid position")
	}
}
