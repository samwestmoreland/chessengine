package tables

import (
	"bytes"
	"testing"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

var testCases = map[sq.Square]uint64{
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

func TestMaskKingAttacks(t *testing.T) {
	for square, expected := range testCases {
		actual := maskKingAttacks(square)
		if uint64(actual) != expected {
			var buf bytes.Buffer

			bb.PrintBoard(actual, &buf)
			t.Error(buf.String())
			t.Errorf("Getting king attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}

func TestPopulateKingAttackTables(t *testing.T) {
	kingAttacks := populateKingAttackTables()

	for square, expected := range testCases {
		actual := kingAttacks[square]
		if uint64(actual) != expected {
			var buf bytes.Buffer

			bb.PrintBoard(actual, &buf)
			t.Error(buf.String())
			t.Errorf("Checking king attack table for square %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}
