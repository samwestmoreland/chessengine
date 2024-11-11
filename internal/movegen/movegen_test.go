package movegen_test

import (
	"os"
	"testing"

	"github.com/samwestmoreland/chessengine/internal/movegen"
	"github.com/samwestmoreland/chessengine/internal/position"
)

func TestMain(m *testing.M) {
	if err := movegen.Initialise(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestGetLegalMoves(t *testing.T) {
	tests := []struct {
		fen      string
		numMoves int
	}{
		{"3n1n2/4P3/8/8/8/8/8/8 w - - 0 1", 2},
		{"8/K7/8/P7/8/3k4/2p5/8 b - - 0 55", 7},
	}

	for _, test := range tests {
		pos, err := position.NewPositionFromFEN(test.fen)
		if err != nil {
			t.Error(err)
		}

		moves := movegen.GetLegalMoves(pos)

		if len(moves) != test.numMoves {
			pos.Print(os.Stderr)
			t.Errorf("Expected %d moves, got %d", test.numMoves, len(moves))
		}
	}
}
