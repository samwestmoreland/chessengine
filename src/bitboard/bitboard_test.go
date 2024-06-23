package bitboard

import (
	"testing"
)

func TestGetBit(t *testing.T) {
	board := setBit(0, e2)
	board = setBit(board, e8)

	if getBit(board, 0) {
		printBoard(board)
		t.Error("Expected false, got true")
	}

	if !getBit(board, e2) {
		printBoard(board)
		t.Error("Expected true, got false")
	}

	if !getBit(board, e8) {
		printBoard(board)
		t.Error("Expected true, got false")
	}

	if getBit(board, f5) {
		printBoard(board)
		t.Error("Expected false, got true")
	}
}
