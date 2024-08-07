package tables

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/bitboard"
	sq "github.com/samwestmoreland/chessengine/src/squares"
)

var rookTestCases = map[int]uint64{
	sq.E4: 4521664529305600,  // central
	sq.G4: 18085034619584512, // g-file
	sq.B5: 565159647117824,   // b-file
	sq.H4: 541165879296,      // h-file
	sq.A6: 8257536,           // a-file
	sq.E1: 4521260802379776,  // 1st rank
	sq.D7: 2260630401218048,  // 7th rank
	sq.H8: 0,                 // corner
	sq.A1: 0,                 // another corner
	sq.H2: 35465847065542656, // another corner
}

func TestComputeRookAttacks(t *testing.T) {
	for square, expected := range rookTestCases {
		actual := computeRookAttacks(square)
		if actual != uint64(expected) {
			bitboard.PrintBoard(actual)
			t.Errorf("Computing rook attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}
