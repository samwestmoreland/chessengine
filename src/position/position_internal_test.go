package position

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

func TestGetPositionFromFEN(t *testing.T) {
	fen, err := ParseFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	if err != nil {
		t.Fatalf("Error in ParseFEN: %s", err)
	}

	pos := getPositionFromFEN(fen)
	if pos == nil {
		t.Fatal("Error in GetPositionFromFEN")
	}

	// Print the position
	t.Log(pos.String())

	// Check that there is a pawn on e4
	e4, err := board.NewSquare("e4")
	if err != nil {
		t.Fatalf("error initialising square: %s", err)
	}

	p := pos.White[e4]
	if p == nil {
		t.Fatal("piece should not be nil")
	}

	if p.Type() != piece.PawnType {
		t.Fatal("error in GetPositionFromFEN")
	}
}

func TestPrintPosition(t *testing.T) {
	e4, _ := board.NewSquare("e4")
	g3, _ := board.NewSquare("g3")
	whiteKing := NewKing(e4, board.White)
	blackBishop := NewBishop(g3, board.Black)
	pos := NewPosition(board.White, []Piece{whiteKing, blackBishop})

	output := pos.String()

	expectedOutput := ". . . . . . . . \n" +
		". . . . . . . . \n" +
		". . . . . . . . \n" +
		". . . . . . . . \n" +
		". . . . K . . . \n" +
		". . . . . . b . \n" +
		". . . . . . . . \n" +
		". . . . . . . . \n"

	if output != expectedOutput {
		t.Fatalf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

func TestSquareIsOccupied(t *testing.T) {
	a2 := board.NewSquareOrPanic("a2")
	pawn1 := NewPawn(a2, board.White)

	pos := NewPosition(board.White, []Piece{pawn1})

	if occ, col := pos.squareIsOccupied(a2); !occ || col != board.White {
		t.Fatalf("Expected square %v to be occupied by white piece", a2)
	}

	if occ, _ := pos.squareIsOccupied(board.NewSquareOrPanic("a3")); occ {
		t.Fatalf("Expected square %v to be unoccupied", a2)
	}
}

func TestGetAllMovesConcurrent(t *testing.T) {
	e4, _ := board.NewSquare("e4")
	g3, _ := board.NewSquare("g3")
	whiteKing := NewKing(e4, board.White)
	whiteBishop := NewBishop(g3, board.White)
	pos := NewPosition(board.White, []Piece{whiteKing, whiteBishop})

	movs, err := pos.GetAllMovesConcurrent(board.White)
	if err != nil {
		t.Fatalf("Error in GetAllMovesConcurrent: %s", err)
	}

	expectedMoves := []moves.Move{}

	expectedSquaresForKing := []string{"e5", "f5", "f4", "f3", "e3", "d3", "d4", "d5"}
	for _, sq := range expectedSquaresForKing {
		s, _ := board.NewSquare(sq)
		expectedMoves = append(expectedMoves, moves.NewMove(e4, s, piece.KingType))
	}

	expectedSquaresForBishop := []string{"h4", "h2", "f2", "e1", "f4", "e5", "d6", "c7", "b8"}
	for _, sq := range expectedSquaresForBishop {
		s, _ := board.NewSquare(sq)
		expectedMoves = append(expectedMoves, moves.NewMove(g3, s, piece.BishopType))
	}

	if len(movs) != len(expectedMoves) {
		t.Log(pos.String())
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(movs))
	}

	if equal := moves.MoveListsEqual(movs, expectedMoves); !equal {
		t.Log(pos.String())
		t.Fatalf("Expected moves %v, got %v", expectedMoves, movs)
	}
}

func TestGetAllMovesConcurrentRealGame(t *testing.T) {
	fen, err := ParseFEN("r1r3k1/4bppp/p3p3/2p1P3/1pP5/3PB2P/P1P2PP1/1R3RK1 w - - 0 17")
	if err != nil {
		t.Fatalf("Error in ParseFEN: %s", err)
	}

	pos := NewPositionFromFEN(fen)

	movs, err := pos.GetAllMovesConcurrent(pos.GetTurn())
	if err != nil {
		t.Fatalf("Error in GetAllMovesConcurrent: %s", err)
	}

	allExpectedMoves := map[piece.Type]map[string][]string{
		piece.PawnType: {
			"a2": {"a3", "a4"},
			"c2": {"c3"},
			"d3": {"d4"},
			"f2": {"f3", "f4"},
			"g2": {"g3", "g4"},
			"h3": {"h4"},
		},
		piece.RookType: {
			"b1": {"b2", "b3", "a1", "c1", "d1", "e1"},
			"f1": {"c1", "d1", "e1"},
		},
		piece.BishopType: {
			"e3": {"c1", "d2", "d4", "f4", "g5", "h6"},
		},
		piece.KingType: {
			"g1": {"h1", "h2"},
		},
	}

	expectedMoves := []moves.Move{}

	for pieceType, expectedMovesForPieceType := range allExpectedMoves {
		for fromSquare, toSquares := range expectedMovesForPieceType {
			sq := board.NewSquareOrPanic(fromSquare)

			for _, toSquare := range toSquares {
				toSq := board.NewSquareOrPanic(toSquare)
				m := moves.NewMove(sq, toSq, pieceType)
				expectedMoves = append(expectedMoves, m)
			}
		}
	}

	if equal := moves.MoveListsEqual(movs, expectedMoves); !equal {
		t.Logf("\n%v", pos)
		t.Fatalf("Expected moves\n%v\ngot\n%v", expectedMoves, movs)
	}
}
