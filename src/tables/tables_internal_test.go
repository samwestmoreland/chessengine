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

func TestComputePawnAttacksWhiteCentral(t *testing.T) {
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

func TestComputePawnAttacksBlackCentral(t *testing.T) {
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

func TestComputePawnAttacksWhiteFlanks(t *testing.T) {
	attackedSquares := computePawnAttacks(bitboard.A2, 0)

	if !bitboard.GetBit(attackedSquares, bitboard.B3) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 2199023255552 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 2199023255552, got ", attackedSquares)
	}

	// reset
	attackedSquares = 0

	attackedSquares = computePawnAttacks(bitboard.H7, 0)

	if !bitboard.GetBit(attackedSquares, bitboard.G8) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 64 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 64, got ", attackedSquares)
	}
}

func TestComputePawnAttacksBlackFlanks(t *testing.T) {
	attackedSquares := computePawnAttacks(bitboard.A7, 1)

	if !bitboard.GetBit(attackedSquares, bitboard.B6) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 131072 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 131072, got ", attackedSquares)
	}

	// reset
	attackedSquares = 0

	attackedSquares = computePawnAttacks(bitboard.H2, 1)

	if !bitboard.GetBit(attackedSquares, bitboard.G1) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 4611686018427387904 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 4611686018427387904, got ", attackedSquares)
	}
}
