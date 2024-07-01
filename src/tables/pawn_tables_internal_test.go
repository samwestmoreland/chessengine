package tables

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/bitboard"
	wb "github.com/samwestmoreland/chessengine/src/colours"
	sq "github.com/samwestmoreland/chessengine/src/squares"
)

func TestIsAFile(t *testing.T) {
	if isAFile(sq.E2) {
		t.Error("Expected false, got true")
	}

	if !isAFile(sq.A3) {
		t.Error("Expected true, got false")
	}
}

func TestIsHFile(t *testing.T) {
	if isHFile(sq.E2) {
		t.Error("Expected false, got true")
	}

	if isHFile(sq.A3) {
		t.Error("Expected false, got true")
	}

	if !isHFile(sq.H8) {
		t.Error("Expected true, got false")
	}
}

func TestComputePawnAttacksWhiteCentral(t *testing.T) {
	attackedSquares := computePawnAttacks(wb.White, sq.E2)

	if !bitboard.GetBit(attackedSquares, sq.D3) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if !bitboard.GetBit(attackedSquares, sq.F3) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 43980465111040 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 43980465111040, got", attackedSquares)
	}
}

func TestComputePawnAttacksBlackCentral(t *testing.T) {
	attackedSquares := computePawnAttacks(wb.Black, sq.C7)

	if !bitboard.GetBit(attackedSquares, sq.B6) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if !bitboard.GetBit(attackedSquares, sq.D6) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 655360 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 655360, got", attackedSquares)
	}
}

func TestComputePawnAttacksWhiteFlanks(t *testing.T) {
	attackedSquares := computePawnAttacks(wb.White, sq.A2)

	if !bitboard.GetBit(attackedSquares, sq.B3) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 2199023255552 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 2199023255552, got", attackedSquares)
	}

	attackedSquares = computePawnAttacks(wb.White, sq.H7)

	if !bitboard.GetBit(attackedSquares, sq.G8) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 64 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 64, got", attackedSquares)
	}
}

func TestComputePawnAttacksBlackFlanks(t *testing.T) {
	attackedSquares := computePawnAttacks(wb.Black, sq.A7)

	if !bitboard.GetBit(attackedSquares, sq.B6) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 131072 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 131072, got", attackedSquares)
	}

	attackedSquares = computePawnAttacks(wb.Black, sq.H2)

	if !bitboard.GetBit(attackedSquares, sq.G1) {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected false, got true")
	}

	if attackedSquares != 4611686018427387904 {
		bitboard.PrintBoard(attackedSquares)
		t.Error("Expected 4611686018427387904, got", attackedSquares)
	}
}

func TestPopulatePawnAttackTables(t *testing.T) {
	PopulatePawnAttackTables()

	if pawnAttacks[wb.White][sq.E2] != 43980465111040 {
		bitboard.PrintBoard(pawnAttacks[wb.White][sq.E2])
		t.Error("Expected 43980465111040, got", pawnAttacks[wb.White][sq.E2])
	}

	if pawnAttacks[wb.Black][sq.F3] != 22517998136852480 {
		bitboard.PrintBoard(pawnAttacks[wb.Black][sq.F3])
		t.Error("Expected 22517998136852480, got", pawnAttacks[wb.Black][sq.F3])
	}

	if pawnAttacks[wb.White][sq.A2] != 2199023255552 {
		bitboard.PrintBoard(pawnAttacks[wb.White][sq.A2])
		t.Error("Expected 2199023255552, got", pawnAttacks[wb.White][sq.A2])
	}

	if pawnAttacks[wb.Black][sq.A2] != 144115188075855872 {
		bitboard.PrintBoard(pawnAttacks[wb.Black][sq.A2])
		t.Error("Expected 144115188075855872, got", pawnAttacks[wb.Black][sq.A2])
	}

	if pawnAttacks[wb.White][sq.H8] != 0 {
		bitboard.PrintBoard(pawnAttacks[wb.White][sq.H8])
		t.Error("Expected 0, got", pawnAttacks[wb.White][sq.H8])
	}

	if pawnAttacks[wb.Black][sq.H7] != 4194304 {
		bitboard.PrintBoard(pawnAttacks[wb.Black][sq.H7])
		t.Error("Expected 4194304, got", pawnAttacks[wb.Black][sq.H7])
	}
}
