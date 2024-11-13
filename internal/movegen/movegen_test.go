package movegen_test

import (
	"os"
	"testing"

	"github.com/samwestmoreland/chessengine/internal/movegen"
	"github.com/samwestmoreland/chessengine/internal/position"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
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
			numMoves: 11, // this test fails because king moves not implemented
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

func TestSquareIsAttacked(t *testing.T) {
	tests := []struct {
		name           string
		fen            string
		square         int
		whiteAttacking bool
		attacked       bool
	}{
		{
			name:           "white pawn on e4 attacking d5",
			fen:            "8/8/8/8/4P3/8/8/8 w - - 0 1",
			square:         sq.D5,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "white pawn on e4 attacking f5",
			fen:            "8/8/8/8/4P3/8/8/8 w - - 0 1",
			square:         sq.F5,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "square not attacked by pawn",
			fen:            "8/8/8/8/4P3/8/8/8 w - - 0 1",
			square:         sq.E5,
			whiteAttacking: true,
			attacked:       false,
		},
		{
			name:           "white pawn attacking black king",
			fen:            "8/8/2k5/3P4/8/8/8/8 w - - 0 1",
			square:         sq.C6,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "white pawn attacking black king on 8th rank",
			fen:            "2k5/1P6/8/8/8/8/8/8 w - - 0 1",
			square:         sq.C8,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "black pawn attacking white pawn",
			fen:            "2k5/8/8/4p3/5P2/8/8/1K6 b - - 0 1",
			square:         sq.F4,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "white knight attacking black pawn",
			fen:            "2k5/8/8/4p3/8/3N4/8/1K6 w - - 0 1",
			square:         sq.F4,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "square not attacked by white knight",
			fen:            "2k5/8/8/4p3/8/3N4/8/1K6 w - - 0 1",
			square:         sq.A7,
			whiteAttacking: true,
			attacked:       false,
		},
		{
			name:           "black knight attacking white king",
			fen:            "2k5/8/8/4p3/8/n2N4/8/1K6 b - - 0 1",
			square:         sq.B1,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "backwards knight attack",
			fen:            "2k5/8/8/4p3/8/n2N4/8/1K6 b - - 0 1",
			square:         sq.B5,
			whiteAttacking: false,
			attacked:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, err := position.NewPositionFromFEN(tt.fen)
			if err != nil {
				t.Fatalf("failed to create position: %v", err)
			}

			isAttacked := movegen.SquareAttacked(pos, tt.square, tt.whiteAttacking)

			if isAttacked != tt.attacked {
				pos.Print(os.Stderr)
				t.Errorf("got %t, want %t", isAttacked, tt.attacked)
			}
		})
	}
}
