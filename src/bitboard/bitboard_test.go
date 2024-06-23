package bitboard

import (
	"testing"
)

func TestGetBit(t *testing.T) {
	var board uint64 = 8 // 1000

	if getBit(board, a8) {
		printBoard(board)
		t.Error("Expected false, got true")
	}

	if !getBit(board, d8) {
		printBoard(board)
		t.Error("Expected true, got false")
	}
}

func TestSetBit(t *testing.T) {
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

func TestClearBit(t *testing.T) {
	board := setBit(0, e2)
	board = setBit(board, e8)

	if !getBit(board, e2) {
		printBoard(board)
		t.Error("Expected true, got false")
	}

	board = clearBit(board, e2)

	if getBit(board, e2) {
		printBoard(board)
		t.Error("Expected false, got true")
	}
}
