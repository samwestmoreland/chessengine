package tables

import (
	"testing"

	"github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

func TestMaskKnightAttacks(t *testing.T) {
	testCases := map[int]uint64{
		sq.E4: 11333767002587136,   // central
		sq.G4: 45053588738670592,   // g-file
		sq.B5: 5531918402816,       // b-file
		sq.H4: 18049583422636032,   // h-file
		sq.A6: 8657044482,          // a-file
		sq.E1: 19184278881435648,   // 1st rank
		sq.D7: 337772578,           // 7th rank
		sq.H8: 4202496,             // corner
		sq.A1: 1128098930098176,    // another corner
		sq.H2: 2305878468463689728, // another corner
	}

	for square, expected := range testCases {
		actual := maskKnightAttacks(square)
		if uint64(actual) != expected {
			bitboard.PrintBoard(actual)
			t.Errorf("Getting knight attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}