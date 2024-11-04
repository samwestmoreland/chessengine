package tables

import (
	"fmt"
	"testing"

	"github.com/samwestmoreland/chessengine/src/bitboard"
	sq "github.com/samwestmoreland/chessengine/src/squares"
)

var rookTestCases = map[int]uint64{
	sq.E4: 4521664529305600,    // central
	sq.G4: 18085034619584512,   // g-file
	sq.B5: 565159647117824,     // b-file
	sq.H4: 36170077829103616,   // h-file
	sq.A6: 282578808340736,     // a-file
	sq.E1: 7930856604974452736, // 1st rank
	sq.D7: 2260630401218048,    // 7th rank
	sq.H8: 36170086419038334,   // corner
	sq.A1: 9079539427579068672, // another corner
	sq.H2: 35607136465616896,   // another corner
}

func TestMaskRookAttacks(t *testing.T) {
	for square, expected := range rookTestCases {
		actual := maskRookAttacks(square)
		if actual != expected {
			fmt.Println("Square", sq.Stringify(square))
			fmt.Println("Got")
			bitboard.PrintBoard(actual)
			fmt.Println("Expected")
			bitboard.PrintBoard(expected)
			fmt.Println("")
			t.Errorf("Computing rook attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}

func TestRookAttacksOnTheFly(t *testing.T) {
	var blockers uint64
	blockers = bitboard.SetBit(blockers, sq.D7)
	blockers = bitboard.SetBit(blockers, sq.D3)
	blockers = bitboard.SetBit(blockers, sq.F4)
	blockers = bitboard.SetBit(blockers, sq.B4)
	blockers = bitboard.SetBit(blockers, sq.A4)

	fmt.Println("Blockers:")
	bitboard.PrintBoard(blockers)

	fmt.Println("Rook attacks on the fly:")
	bitboard.PrintBoard(rookAttacksOnTheFly(sq.D4, blockers))

	t.Errorf("Artificial failure")
}

func TestSetOccupancy(t *testing.T) {
	rookAttacks := maskRookAttacks(sq.D4)

	fmt.Println("Rook attacks:")
	bitboard.PrintBoard(rookAttacks)

	fmt.Println("Set occupancy:")
	bitboard.PrintBoard(bitboard.SetOccupancy(9, rookAttacks))

	t.Errorf("Artificial failure")
}
