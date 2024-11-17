package tables

import (
	"bytes"
	"testing"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
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

func TestMaskPawnAttacksWhiteCentral(t *testing.T) {
	attackedSquares := maskPawnAttacks(0, sq.E2)

	if !bb.GetBit(attackedSquares, sq.D3) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if !bb.GetBit(attackedSquares, sq.F3) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if attackedSquares != 43980465111040 {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected 43980465111040, got", attackedSquares)
	}
}

func TestMaskPawnAttacksBlackCentral(t *testing.T) {
	attackedSquares := maskPawnAttacks(1, sq.C7)

	if !bb.GetBit(attackedSquares, sq.B6) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if !bb.GetBit(attackedSquares, sq.D6) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if attackedSquares != 655360 {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected 655360, got", attackedSquares)
	}
}

func TestMaskPawnAttacksWhiteFlanks(t *testing.T) {
	attackedSquares := maskPawnAttacks(0, sq.A2)

	if !bb.GetBit(attackedSquares, sq.B3) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if attackedSquares != 2199023255552 {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected 2199023255552, got", attackedSquares)
	}

	attackedSquares = maskPawnAttacks(0, sq.H7)

	if !bb.GetBit(attackedSquares, sq.G8) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if attackedSquares != 64 {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected 64, got", attackedSquares)
	}
}

func TestMaskPawnAttacksBlackFlanks(t *testing.T) {
	attackedSquares := maskPawnAttacks(1, sq.A7)

	if !bb.GetBit(attackedSquares, sq.B6) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if attackedSquares != 131072 {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected 131072, got", attackedSquares)
	}

	attackedSquares = maskPawnAttacks(1, sq.H2)

	if !bb.GetBit(attackedSquares, sq.G1) {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected false, got true")
	}

	if attackedSquares != 4611686018427387904 {
		var buf bytes.Buffer

		bb.PrintBoard(attackedSquares, &buf)
		t.Error(buf.String())
		t.Error("Expected 4611686018427387904, got", attackedSquares)
	}
}

func TestPopulatePawnAttackTables(t *testing.T) {
	pawnAttacks := populatePawnAttackTables()

	if pawnAttacks[0][sq.E2] != 43980465111040 {
		var buf bytes.Buffer

		bb.PrintBoard(pawnAttacks[0][sq.E2], &buf)
		t.Error(buf.String())
		t.Error("Expected 43980465111040, got", pawnAttacks[0][sq.E2])
	}

	if pawnAttacks[1][sq.F3] != 22517998136852480 {
		var buf bytes.Buffer

		bb.PrintBoard(pawnAttacks[1][sq.F3], &buf)
		t.Error(buf.String())
		t.Error("Expected 22517998136852480, got", pawnAttacks[1][sq.F3])
	}

	if pawnAttacks[0][sq.A2] != 2199023255552 {
		var buf bytes.Buffer

		bb.PrintBoard(pawnAttacks[0][sq.A2], &buf)
		t.Error(buf.String())
		t.Error("Expected 2199023255552, got", pawnAttacks[0][sq.A2])
	}

	if pawnAttacks[1][sq.A2] != 144115188075855872 {
		var buf bytes.Buffer

		bb.PrintBoard(pawnAttacks[1][sq.A2], &buf)
		t.Error(buf.String())
		t.Error("Expected 144115188075855872, got", pawnAttacks[1][sq.A2])
	}

	if pawnAttacks[0][sq.H8] != 0 {
		var buf bytes.Buffer

		bb.PrintBoard(pawnAttacks[0][sq.H8], &buf)
		t.Error(buf.String())
		t.Error("Expected 0, got", pawnAttacks[0][sq.H8])
	}

	if pawnAttacks[1][sq.H7] != 4194304 {
		var buf bytes.Buffer

		bb.PrintBoard(pawnAttacks[1][sq.H7], &buf)
		t.Error(buf.String())
		t.Error("Expected 4194304, got", pawnAttacks[1][sq.H7])
	}
}
