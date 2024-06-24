package tables

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/bitboard"
)

func TestIsAFile(t *testing.T) {
	if isAFile(bitboard.E2) {
		t.Error("Expected false, got true")
	}

	if !isAFile(bitboard.A3) {
		t.Error("Expected true, got false")
	}
}

func TestIsHFile(t *testing.T) {
	if isHFile(bitboard.E2) {
		t.Error("Expected false, got true")
	}

	if isHFile(bitboard.A3) {
		t.Error("Expected false, got true")
	}

	if !isHFile(bitboard.H8) {
		t.Error("Expected true, got false")
	}
}

func TestComputePawnAttacksWhite(t *testing.T) {
	attackedSquares := computePawnAttacks(bitboard.E2, 0)

	if !bitboard.GetBit(attackedSquares, bitboard.D3) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if !bitboard.GetBit(attackedSquares, bitboard.F3) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 43980465111040 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 43980465111040, got ", attackedSquares)
	}
}

func TestComputePawnAttacksBlack(t *testing.T) {
	attackedSquares := computePawnAttacks(bitboard.C7, 1)

	if !bitboard.GetBit(attackedSquares, bitboard.B6) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if !bitboard.GetBit(attackedSquares, bitboard.D6) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 655360 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 655360, got ", attackedSquares)
	}
}
