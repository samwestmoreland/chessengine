package movegen_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/samwestmoreland/chessengine/internal/move"
	"github.com/samwestmoreland/chessengine/internal/movegen"
	"github.com/samwestmoreland/chessengine/internal/piece"
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
	t.Parallel()

	tests := []struct {
		name     string // Consider adding test case names
		fen      string
		numMoves int
	}{
		{
			name:     "starting position",
			fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			numMoves: 20,
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
		{
			name:     "white to castle king side",
			fen:      "1k6/4b3/8/8/8/8/8/4K2R w K - 0 1",
			numMoves: 15,
		},
		{
			name:     "white to castle queen side",
			fen:      "1k6/4b3/8/8/8/8/8/R3K3 w Q - 0 1",
			numMoves: 16,
		},
		{
			name:     "white knight moves",
			fen:      "6k1/8/2P5/8/3N4/8/8/8 w - - 0 1",
			numMoves: 8,
		},
		{
			name:     "white bishop moves",
			fen:      "6k1/8/2P5/8/4B3/8/2P5/8 w - - 0 1",
			numMoves: 11,
		},
		{
			name:     "white king moves",
			fen:      "6k1/8/2P5/8/4B3/4K3/2P5/8 w - - 0 1",
			numMoves: 18,
		},
		{
			name:     "white queen moves",
			fen:      "1k6/8/6r1/8/4Q3/8/2K5/8 w - - 0 1",
			numMoves: 32,
		},
		{
			name:     "van't kruijs opening black to move",
			fen:      "rnbq1rk1/ppp1ppbp/3p1np1/8/3P4/4P1P1/PPPNNPBP/R1BQK2R b KQ - 1 6",
			numMoves: 32,
		},
		{
			name:     "van't kruijs opening white to move",
			fen:      "rnbq1rk1/ppp2pbp/3p1np1/4p3/3P4/4P1P1/PPPNNPBP/R1BQK2R w KQ - 0 7",
			numMoves: 35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pos, err := position.NewPositionFromFEN(tt.fen)
			if err != nil {
				t.Fatalf("failed to create position: %v", err)
			}

			moves := movegen.GetLegalMoves(pos)

			if len(moves) != tt.numMoves {
				var buf bytes.Buffer

				buf.WriteString("\n")
				pos.Print(&buf)
				t.Error(buf.String())
				t.Errorf("got %d moves, want %d", len(moves), tt.numMoves)
				t.Errorf("moves generated:")

				for _, move := range moves {
					t.Errorf("  %s", move.String())
				}
			}
		})
	}
}

func TestGetLegalMovesReturnsCorrectMoves(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		fen           string
		expectedMoves map[move.Move]struct{}
	}{
		{
			name: "evan's gambit white to move",
			fen:  "r1bqk1nr/pppp1ppp/2n5/b3p3/2BPP3/2P2N2/P4PPP/RNBQK2R b KQkq - 0 3",
			expectedMoves: map[move.Move]struct{}{
				// Pawns
				move.NewMove().From(sq.A7).To(sq.A6).Piece(piece.Bp).Build():              {},
				move.NewMove().From(sq.B7).To(sq.B6).Piece(piece.Bp).Build():              {},
				move.NewMove().From(sq.B7).To(sq.B5).Piece(piece.Bp).DoublePush().Build(): {},
				move.NewMove().From(sq.D7).To(sq.D6).Piece(piece.Bp).Build():              {},
				move.NewMove().From(sq.D7).To(sq.D5).Piece(piece.Bp).DoublePush().Build(): {},
				move.NewMove().From(sq.F7).To(sq.F6).Piece(piece.Bp).Build():              {},
				move.NewMove().From(sq.F7).To(sq.F5).Piece(piece.Bp).DoublePush().Build(): {},
				move.NewMove().From(sq.G7).To(sq.G6).Piece(piece.Bp).Build():              {},
				move.NewMove().From(sq.G7).To(sq.G5).Piece(piece.Bp).DoublePush().Build(): {},
				move.NewMove().From(sq.H7).To(sq.H6).Piece(piece.Bp).Build():              {},
				move.NewMove().From(sq.H7).To(sq.H5).Piece(piece.Bp).DoublePush().Build(): {},
				move.NewMove().From(sq.E5).To(sq.D4).Piece(piece.Bp).Capture().Build():    {},
				// Rooks
				move.NewMove().From(sq.A8).To(sq.B8).Piece(piece.Br).Build(): {},
				// Bishops
				move.NewMove().From(sq.A5).To(sq.B6).Piece(piece.Bb).Build():           {},
				move.NewMove().From(sq.A5).To(sq.B4).Piece(piece.Bb).Build():           {},
				move.NewMove().From(sq.A5).To(sq.C3).Piece(piece.Bb).Capture().Build(): {},
				// Knights
				move.NewMove().From(sq.C6).To(sq.B8).Piece(piece.Bn).Build():           {},
				move.NewMove().From(sq.C6).To(sq.E7).Piece(piece.Bn).Build():           {},
				move.NewMove().From(sq.C6).To(sq.D4).Piece(piece.Bn).Capture().Build(): {},
				move.NewMove().From(sq.C6).To(sq.B4).Piece(piece.Bn).Build():           {},
				move.NewMove().From(sq.G8).To(sq.H6).Piece(piece.Bn).Build():           {},
				move.NewMove().From(sq.G8).To(sq.F6).Piece(piece.Bn).Build():           {},
				move.NewMove().From(sq.G8).To(sq.E7).Piece(piece.Bn).Build():           {},
				// Queen
				move.NewMove().From(sq.D8).To(sq.E7).Piece(piece.Bq).Build(): {},
				move.NewMove().From(sq.D8).To(sq.F6).Piece(piece.Bq).Build(): {},
				move.NewMove().From(sq.D8).To(sq.G5).Piece(piece.Bq).Build(): {},
				move.NewMove().From(sq.D8).To(sq.H4).Piece(piece.Bq).Build(): {},
				// King
				move.NewMove().From(sq.E8).To(sq.E7).Piece(piece.Bk).Build(): {},
				move.NewMove().From(sq.E8).To(sq.F8).Piece(piece.Bk).Build(): {},
			},
		},
		{
			name: "scandinavian hikaru vs magnus",
			fen:  "4r2k/5Rp1/bppp3p/4b3/4P3/2N5/PPP3PP/3R2K1 w - - 0 23",
			expectedMoves: map[move.Move]struct{}{
				// Pawns
				move.NewMove().From(sq.A2).To(sq.A3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.A2).To(sq.A4).Piece(piece.Wp).DoublePush().Build(): {},
				move.NewMove().From(sq.B2).To(sq.B3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.B2).To(sq.B4).Piece(piece.Wp).DoublePush().Build(): {},
				move.NewMove().From(sq.G2).To(sq.G3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.G2).To(sq.G4).Piece(piece.Wp).DoublePush().Build(): {},
				move.NewMove().From(sq.H2).To(sq.H3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.H2).To(sq.H4).Piece(piece.Wp).DoublePush().Build(): {},
				// Rooks
				move.NewMove().From(sq.D1).To(sq.D2).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.D3).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.D4).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.D5).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.D6).Piece(piece.Wr).Capture().Build(): {},
				move.NewMove().From(sq.D1).To(sq.A1).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.B1).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.C1).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.E1).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.D1).To(sq.F1).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.F8).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.A7).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.B7).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.C7).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.D7).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.E7).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.F6).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.F5).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.F4).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.F3).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.F2).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.F1).Piece(piece.Wr).Build():           {},
				move.NewMove().From(sq.F7).To(sq.G7).Piece(piece.Wr).Capture().Build(): {},
				// Knights
				move.NewMove().From(sq.C3).To(sq.B1).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C3).To(sq.A4).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C3).To(sq.B5).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C3).To(sq.D5).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C3).To(sq.E2).Piece(piece.Wn).Build(): {},
				// King
				move.NewMove().From(sq.G1).To(sq.F1).Piece(piece.Wk).Build(): {},
				move.NewMove().From(sq.G1).To(sq.F2).Piece(piece.Wk).Build(): {},
				move.NewMove().From(sq.G1).To(sq.H1).Piece(piece.Wk).Build(): {},
			},
		},
		{
			name: "hikaru kasparov scotch gambit",
			fen:  "1rb2rk1/ppN1nppp/3p4/2b2q1B/5Pn1/2P1B3/PP1N2PP/R2Q1R1K w - - 1 16",
			expectedMoves: map[move.Move]struct{}{
				// Pawns
				move.NewMove().From(sq.A2).To(sq.A3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.A2).To(sq.A4).Piece(piece.Wp).DoublePush().Build(): {},
				move.NewMove().From(sq.B2).To(sq.B3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.B2).To(sq.B4).Piece(piece.Wp).DoublePush().Build(): {},
				move.NewMove().From(sq.C3).To(sq.C4).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.G2).To(sq.G3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.H2).To(sq.H3).Piece(piece.Wp).Build():              {},
				move.NewMove().From(sq.H2).To(sq.H4).Piece(piece.Wp).DoublePush().Build(): {},
				// Rooks
				move.NewMove().From(sq.A1).To(sq.B1).Piece(piece.Wr).Build(): {},
				move.NewMove().From(sq.A1).To(sq.C1).Piece(piece.Wr).Build(): {},
				move.NewMove().From(sq.F1).To(sq.F2).Piece(piece.Wr).Build(): {},
				move.NewMove().From(sq.F1).To(sq.G1).Piece(piece.Wr).Build(): {},
				move.NewMove().From(sq.F1).To(sq.F3).Piece(piece.Wr).Build(): {},
				move.NewMove().From(sq.F1).To(sq.E1).Piece(piece.Wr).Build(): {},
				// Bishops
				move.NewMove().From(sq.E3).To(sq.D4).Piece(piece.Wb).Build():           {},
				move.NewMove().From(sq.E3).To(sq.C5).Piece(piece.Wb).Capture().Build(): {},
				move.NewMove().From(sq.E3).To(sq.F2).Piece(piece.Wb).Build():           {},
				move.NewMove().From(sq.E3).To(sq.G1).Piece(piece.Wb).Build():           {},
				move.NewMove().From(sq.H5).To(sq.G6).Piece(piece.Wb).Build():           {},
				move.NewMove().From(sq.H5).To(sq.G4).Piece(piece.Wb).Capture().Build(): {},
				move.NewMove().From(sq.H5).To(sq.F7).Piece(piece.Wb).Capture().Build(): {},
				// Knights
				move.NewMove().From(sq.D2).To(sq.B1).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.D2).To(sq.B3).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.D2).To(sq.C4).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.D2).To(sq.E4).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.D2).To(sq.F3).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C7).To(sq.A8).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C7).To(sq.E8).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C7).To(sq.E6).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C7).To(sq.D5).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C7).To(sq.B5).Piece(piece.Wn).Build(): {},
				move.NewMove().From(sq.C7).To(sq.A6).Piece(piece.Wn).Build(): {},
				// Queen
				move.NewMove().From(sq.D1).To(sq.B1).Piece(piece.Wq).Build():           {},
				move.NewMove().From(sq.D1).To(sq.C1).Piece(piece.Wq).Build():           {},
				move.NewMove().From(sq.D1).To(sq.C2).Piece(piece.Wq).Build():           {},
				move.NewMove().From(sq.D1).To(sq.B3).Piece(piece.Wq).Build():           {},
				move.NewMove().From(sq.D1).To(sq.A4).Piece(piece.Wq).Build():           {},
				move.NewMove().From(sq.D1).To(sq.E2).Piece(piece.Wq).Build():           {},
				move.NewMove().From(sq.D1).To(sq.F3).Piece(piece.Wq).Build():           {},
				move.NewMove().From(sq.D1).To(sq.G4).Piece(piece.Wq).Capture().Build(): {},
				move.NewMove().From(sq.D1).To(sq.E1).Piece(piece.Wq).Build():           {},
				// King
				move.NewMove().From(sq.H1).To(sq.G1).Piece(piece.Wk).Build(): {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pos, err := position.NewPositionFromFEN(tt.fen)
			if err != nil {
				t.Fatalf("failed to create position: %v", err)
			}

			moves := movegen.GetLegalMoves(pos)

			if len(moves) != len(tt.expectedMoves) {
				t.Errorf("moves generated:")

				for _, move := range moves {
					t.Errorf("  %s", move.String())
				}

				t.Errorf("got %d moves, want %d", len(moves), len(tt.expectedMoves))
			}

			for _, move := range moves {
				if _, ok := tt.expectedMoves[move]; !ok {
					t.Errorf("unexpected move: %s; piece: %s; capture: %t; double push: %t",
						move.String(), move.Piece().String(), move.IsCapture(), move.IsDoublePush())
				}
			}
		})
	}
}

func TestSquareIsAttacked(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		fen            string
		square         sq.Square
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
		{
			name:           "white king attacking e4",
			fen:            "1k6/8/8/4r3/3K4/8/8/8 w - - 0 1",
			square:         sq.E4,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "square not attacked by anything",
			fen:            "1k6/8/8/4r3/3K4/8/8/8 w - - 0 1",
			square:         sq.B3,
			whiteAttacking: true,
			attacked:       false,
		},
		{
			name:           "black king attacking a8",
			fen:            "1k6/8/8/4r3/3K4/8/8/8 b - - 0 1",
			square:         sq.A8,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "black king attacking a7",
			fen:            "1k6/8/8/4r3/3K4/8/8/8 b - - 0 1",
			square:         sq.A7,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "black king attacking b7",
			fen:            "1k6/8/8/4r3/3K4/8/8/8 b - - 0 1",
			square:         sq.B7,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "black king attacking c7",
			fen:            "1k6/8/8/4r3/3K4/8/8/8 b - - 0 1",
			square:         sq.C7,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "black king attacking c8",
			fen:            "1k6/8/8/4r3/3K4/8/8/8 b - - 0 1",
			square:         sq.C8,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "white bishop attacking e5",
			fen:            "1k6/8/8/4r3/8/2B5/8/1K6 w - - 0 1",
			square:         sq.E5,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "white bishop not attacking blocked square",
			fen:            "1k6/8/8/4r3/8/2B5/8/1K6 w - - 0 1",
			square:         sq.F6,
			whiteAttacking: true,
			attacked:       false,
		},
		{
			name:           "black bishop not attacking F4",
			fen:            "1k6/2b5/8/4r3/8/2B5/8/1K6 b - - 0 1",
			square:         sq.F4,
			whiteAttacking: false,
			attacked:       false,
		},
		{
			name:           "white rook attack 3rd rank",
			fen:            "1k6/8/8/4r3/8/1R6/8/1K6 w - - 0 1",
			square:         sq.H3,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "white rook 3rd rank with blocker",
			fen:            "1k6/8/8/8/8/1R2r3/8/1K6 w - - 0 1",
			square:         sq.F3,
			whiteAttacking: true,
			attacked:       false,
		},
		{
			name:           "white rook attacking another rook",
			fen:            "1k6/8/8/8/8/1R2r3/8/1K6 w - - 0 1",
			square:         sq.E3,
			whiteAttacking: true,
			attacked:       true,
		},
		{
			name:           "black rook not under attack because it is black to move",
			fen:            "1k6/8/8/8/8/1R2r3/8/1K6 b - - 0 1",
			square:         sq.E3,
			whiteAttacking: false,
			attacked:       false,
		},
		{
			name:           "black rook attacking e file",
			fen:            "1k6/8/8/8/8/1R2r3/8/1K6 b - - 0 1",
			square:         sq.E8,
			whiteAttacking: false,
			attacked:       true,
		},
		{
			name:           "black rook attack blocked by pawn of own colour",
			fen:            "1k6/8/8/8/8/1Rp1r3/8/1K6 b - - 0 1",
			square:         sq.B3,
			whiteAttacking: false,
			attacked:       false,
		},
		{
			name:           "black rook attacked by white queen",
			fen:            "1k6/8/8/8/3r4/8/5Q2/1K6 w - - 0 1",
			square:         sq.D4,
			whiteAttacking: true,
			attacked:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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
