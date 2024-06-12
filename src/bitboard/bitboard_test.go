package bitboard

import (
	"testing"
)

func TestGetBit(t *testing.T) {
	var board uint64 = 8

	if getBit(board, 0) {
		t.Error("Expected false, got true")
	}

	if !getBit(board, 3) {
		print(board)
		t.Error("Expected true, got false")
	}
}
