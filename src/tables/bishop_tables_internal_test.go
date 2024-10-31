package tables

import (
	"fmt"
	"testing"

	"github.com/samwestmoreland/chessengine/src/bitboard"
	sq "github.com/samwestmoreland/chessengine/src/squares"
)

var bishopTestCases = map[int]uint64{
	sq.E4: 19184279556981248, // central
	sq.G4: 4538784537380864,  // g-file
	sq.B5: 4512412900526080,  // b-file
	sq.H4: 9077569074761728,  // h-file
	sq.A6: 4512412933816832,  // a-file
	sq.E1: 11333774449049600, // 1st rank
	sq.D7: 275449643008,      // 7th rank
	sq.H8: 567382630219776,   // corner
	sq.A1: 567382630219776,   // another corner
	sq.H2: 70506452091904,    // another corner
}

func TestMaskBishopAttacks(t *testing.T) {
	for square, expected := range bishopTestCases {
		actual := maskBishopAttacks(square)
		if actual != expected {
			bitboard.PrintBoard(actual)
			t.Errorf("Computing bishop attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}

func TestBishopAttacksOnTheFly(t *testing.T) {
	var blockers uint64
	blockers = bitboard.SetBit(blockers, sq.B6)
	blockers = bitboard.SetBit(blockers, sq.G7)
	blockers = bitboard.SetBit(blockers, sq.E3)
	blockers = bitboard.SetBit(blockers, sq.B2)

	fmt.Println("Blockers:")
	bitboard.PrintBoard(blockers)

	fmt.Println("Bishop attacks on the fly:")
	bitboard.PrintBoard(bishopAttacksOnTheFly(sq.D4, blockers))

	t.Errorf("Artificial failure")
}
