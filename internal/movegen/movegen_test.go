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
		name     string // Consider adding test case names
		fen      string
		numMoves int
	}{
		{
			name:     "starting position",
			fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			numMoves: 16,
		},
		{
			name:     "white pawn promotion with captures",
			fen:      "3n1n2/4P3/8/8/8/8/8/8 w - - 0 1",
			numMoves: 12,
		},
		{
			name:     "white pawn en passant",
			fen:      "8/8/8/1PPp4/8/8/8/8 w - d6 0 1",
			numMoves: 3,
		},
		{
			name:     "single black pawn promoting",
			fen:      "8/8/8/8/8/8/2p5/8 b - - 0 1",
			numMoves: 4,
		},
		{
			name:     "black pawn in middle of board",
			fen:      "8/8/8/5p2/8/8/8/8 b - - 0 1",
			numMoves: 1,
		},
		{
			name:     "black pawns in starting position",
			fen:      "8/2p2p2/8/8/8/8/8/8 b - - 0 1",
			numMoves: 4,
		},
		{
			name:     "black pawn en passant",
			fen:      "8/8/8/1P6/3pP3/8/8/8 b - e3 0 1",
			numMoves: 2,
		},
		{
			name:     "black pawn and king moves",
			fen:      "8/K7/8/P7/8/3k4/2p5/8 b - - 0 1",
			numMoves: 12, // this test fails because king moves not implemented
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, err := position.NewPositionFromFEN(tt.fen)
			if err != nil {
				t.Fatalf("failed to create position: %v", err)
			}

			moves := movegen.GetLegalMoves(pos)

			if len(moves) != tt.numMoves {
				pos.Print(os.Stderr)
				t.Errorf("got %d moves, want %d", len(moves), tt.numMoves)
				t.Errorf("moves generated:")
				for _, move := range moves {
					t.Errorf("  %s", move.String())
				}
			}
		})
	}
}
