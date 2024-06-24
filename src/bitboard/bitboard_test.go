package bitboard

import (
	"testing"
)

func TestGetBit(t *testing.T) {
	var board uint64 = 8 // 1000

	if GetBit(board, A8) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}

	if !GetBit(board, D8) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}
}

func TestSetBit(t *testing.T) {
	board := SetBit(0, E2)
	board = SetBit(board, E8)

	if GetBit(board, 0) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}

	if !GetBit(board, E2) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}

	if !GetBit(board, E8) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}

	if GetBit(board, F5) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}
}

func TestClearBit(t *testing.T) {
	board := SetBit(0, E2)
	board = SetBit(board, E8)

	if !GetBit(board, E2) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}

	board = ClearBit(board, E2)

	if GetBit(board, E2) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}
}
