package position

import (
	"testing"
)

func TestValidPositionCheck(t *testing.T) {
	t.Log("Testing valid position check")
	var pos Position = "r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3"
	if !pos.isValid() {
		t.Error("Expected valid position")
	}
}

func TestInvalidPosition(t *testing.T) {
	t.Log("Testing invalid position check")
	var pos Position = "pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 2 3"
	if pos.isValid() {
		t.Error("Expected invalid position")
	}
}

func TestSplitRows(t *testing.T) {
	t.Log("Testing split rows")
	var pos Position = "r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3"
	rows := pos.splitRows()
	if len(rows) != 8 {
		t.Error("Expected 8 rows")
	}
}
