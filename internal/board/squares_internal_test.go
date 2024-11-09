package board

import (
	"testing"
)

func TestSquares(t *testing.T) {
	if len(Squares) != 64 {
		t.Error("Expected 64 squares, got", len(Squares))
	}
}
