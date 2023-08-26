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

	expectedMoves := moves.MoveList{}

	expectedSquaresForKing := []string{"e5", "f5", "f4", "f3", "e3", "d3", "d4", "d5"}
	for _, sq := range expectedSquaresForKing {
		s, _ := board.NewSquare(sq)
		expectedMoves.AddMove(moves.NewMove(e4, s, piece.KingType, false))
	}

	expectedSquaresForBishop := []string{"h4", "h2", "f2", "e1", "f4", "e5", "d6", "c7", "b8"}
	for _, sq := range expectedSquaresForBishop {
		s, _ := board.NewSquare(sq)
		expectedMoves.AddMove(moves.NewMove(g3, s, piece.BishopType, false))
	}

	if movs.Len() != expectedMoves.Len() {
		t.Log(pos.String())
		t.Fatalf("Expected %d moves, got %d", expectedMoves.Len(), movs.Len())
	}

	if equal := movs.Equals(expectedMoves); !equal {
		t.Log(pos.String())
		t.Fatalf("Expected moves %v, got %v", expectedMoves, movs)
	}
}

func TestGetAllMovesWithCaptures(t *testing.T) {
	fen, err := ParseFEN("8/8/K1k2p2/8/P4bpP/8/7P/8 b - - 1 48")
	if err != nil {
		t.Fatalf("Error in ParseFEN: %s", err)
	}

	pos := NewPositionFromFEN(fen)

	_, err = pos.GetAllMovesConcurrent(pos.GetTurn())
	if err != nil {
		t.Fatalf("Error in GetAllMovesConcurrent: %s", err)
	}

	allExpectedMoves := map[piece.Type]map[string][]map[string]bool{
		piece.PawnType: {
			"f6": {{"f5": false}},
			"g4": {{"g3": false}},
		},
		piece.BishopType: {
			"f4": {
				{"e3": false},
				{"d2": false},
				{"c1": false},
				{"g3": false},
				{"h2": true},
				{"g5": false},
				{"h6": false},
				{"e5": false},
				{"d6": false},
				{"c7": false},
				{"b8": false},
			},
		},
		piece.KingType: {
			"c6": {
				{"c5": false},
				{"d5": false},
				{"d6": false},
				{"d7": false},
				{"c7": false},
			},
		},
	}

	expectedMoves := moves.MoveList{}

	for pieceType, expectedMovesForPieceType := range allExpectedMoves {
		for fromSquare, toSquares := range expectedMovesForPieceType {
			sq := board.NewSquareOrPanic(fromSquare)

			for _, toSquare := range toSquares {
				for toSquareStr, isCapture := range toSquare {
					toSq := board.NewSquareOrPanic(toSquareStr)
					expectedMoves.AddMove(moves.NewMove(sq, toSq, pieceType, isCapture))
				}
			}
		}
	}

}
