package bitboard

import (
	"bytes"
	"testing"

	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

func TestGetBit(t *testing.T) {
	t.Parallel()

	var board Bitboard = 8 // 1000

	if GetBit(board, sq.A8) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if !GetBit(board, sq.D8) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected true, got false")
	}
}

func TestSetBit(t *testing.T) {
	t.Parallel()

	board := SetBit(0, sq.E2)
	board = SetBit(board, sq.E8)

	if GetBit(board, 0) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if !GetBit(board, sq.E2) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected true, got false")
	}

	if !GetBit(board, sq.E8) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected true, got false")
	}

	if GetBit(board, sq.F5) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}
}

func TestSetWholeBoard(t *testing.T) {
	t.Parallel()

	var board Bitboard

	for rank := range 8 {
		for file := range 8 {
			square := sq.Square(byte(rank*8 + file))
			board = SetBit(board, square)
		}
	}

	if board != 18446744073709551615 {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected 18446744073709551615, got ", board)
	}
}

func TestClearBit(t *testing.T) {
	t.Parallel()

	board := SetBit(0, sq.E2)
	board = SetBit(board, sq.E8)

	if !GetBit(board, sq.E2) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected true, got false")
	}

	board = ClearBit(board, sq.E2)

	if GetBit(board, sq.E2) {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}
}

func TestLSBIndex(t *testing.T) {
	t.Parallel()

	board := SetBit(0, sq.E2)
	board = SetBit(board, sq.E8)
	board = SetBit(board, sq.F5)

	index := LSBIndex(board)
	if index != 4 {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected 4, got ", index)
	}

	asSquare := sq.Stringify(index)

	if asSquare != "e8" {
		var buf bytes.Buffer

		PrintBoard(board, &buf)
		t.Error(buf.String())
		t.Error("Expected e8, got ", asSquare)
	}
}

func TestLSBIndexOfZero(t *testing.T) {
	t.Parallel()

	if LSBIndex(0) != sq.NoSquare {
		t.Errorf("Expected %d, got %d", sq.NoSquare, LSBIndex(0))
	}
}
