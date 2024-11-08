package tables

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/bitboard"
	sq "github.com/samwestmoreland/chessengine/src/squares"
)

var testCases = map[int]uint64{
	sq.E4: 61745389371392,       // central
	sq.G4: 246981557485568,      // g-file
	sq.B5: 30149115904,          // b-file
	sq.H4: 211384331665408,      // h-file
	sq.A6: 50463488,             // a-file
	sq.E1: 2898066360212914176,  // 1st rank
	sq.D7: 1840156,              // 7th rank
	sq.H8: 49216,                // corner
	sq.A1: 144959613005987840,   // another corner
	sq.H2: 13853283560024178688, // another corner
}

func TestComputeKingAttacks(t *testing.T) {
	for square, expected := range testCases {
		actual := computeKingAttacks(square)
		if actual != expected {
			bitboard.PrintBoard(actual)
			t.Errorf("Computing king attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}

func TestPopulateKingAttackTables(t *testing.T) {
	kingAttacks := populateKingAttackTables()

	for square, expected := range testCases {
		actual := kingAttacks[square]
		if actual != expected {
			bitboard.PrintBoard(actual)
			t.Errorf("Checking king attack table for square %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}
