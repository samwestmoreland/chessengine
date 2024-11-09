package bitboard

import (
	"testing"

	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

func TestGetBit(t *testing.T) {
	var board Bitboard = 8 // 1000

	if GetBit(board, sq.A8) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}

	if !GetBit(board, sq.D8) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}
}

func TestSetBit(t *testing.T) {
	board := SetBit(0, sq.E2)
	board = SetBit(board, sq.E8)

	if GetBit(board, 0) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}

	if !GetBit(board, sq.E2) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}

	if !GetBit(board, sq.E8) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}

	if GetBit(board, sq.F5) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}
}

func TestSetWholeBoard(t *testing.T) {
	var board Bitboard

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			board = SetBit(board, square)
		}
	}

	if board != 18446744073709551615 {
		PrintBoard(board)
		t.Error("Expected 18446744073709551615, got ", board)
	}
}

func TestClearBit(t *testing.T) {
	board := SetBit(0, sq.E2)
	board = SetBit(board, sq.E8)

	if !GetBit(board, sq.E2) {
		PrintBoard(board)
		t.Error("Expected true, got false")
	}

	board = ClearBit(board, sq.E2)

	if GetBit(board, sq.E2) {
		PrintBoard(board)
		t.Error("Expected false, got true")
	}
}

func TestLSBIndex(t *testing.T) {
	board := SetBit(0, sq.E2)
	board = SetBit(board, sq.E8)
	board = SetBit(board, sq.F5)

	index := LSBIndex(board)
	if index != 4 {
		PrintBoard(board)
		t.Error("Expected 4, got ", index)
	}

	asSquare := sq.Stringify(index)

	if asSquare != "e8" {
		PrintBoard(board)
		t.Error("Expected e8, got ", asSquare)
	}
}

func TestLSBIndexOfZero(t *testing.T) {
	if LSBIndex(0) != -1 {
		t.Error("Expected -1, got ", LSBIndex(0))
	}
}
