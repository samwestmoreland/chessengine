package position

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/piece"
)

func TestGetMovesForKingOnEmptyBoard(t *testing.T) {
	square := board.NewSquareOrPanic("e4")

	whiteKing := NewKing(square, board.White)
	pos := NewPosition(board.White, []Piece{whiteKing})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	expectedSquares := []string{"d5", "e5", "f5", "d4", "f4", "d3", "e3", "f3"}
	expectedMoves := moves.MoveList{}

	for _, sqStr := range expectedSquares {
		sq := board.NewSquareOrPanic(sqStr)
		expectedMoves.AddMove(moves.NewMove(square, sq, piece.KingType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Errorf("expected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Fatalf("expected moves %v, got %v", expectedMoves, mov)
	}
}

func TestGetMovesForKingOnEmptyBoardInCorner(t *testing.T) {
	a1 := board.NewSquareOrPanic("a1")
	whiteKing := NewKing(a1, board.White)

	pos := NewPosition(board.White, []Piece{whiteKing})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	expectedMoves := moves.MoveList{}

	expectedSquares := []string{"b1", "b2", "a2"}
	for _, sqStr := range expectedSquares {
		sq := board.NewSquareOrPanic(sqStr)
		expectedMoves.AddMove(moves.NewMove(a1, sq, piece.KingType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Fatalf("Expected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Fatalf("Expected moves %v, got %v", expectedMoves, mov)
	}
}

func TestGetMovesForKingWhenAnotherPieceOccupiesOneOfThePossibleSquares(t *testing.T) {
	b3 := board.NewSquareOrPanic("b3")
	whiteKing := NewKing(b3, board.White)

	b4 := board.NewSquareOrPanic("b4")
	whiteBishop := NewBishop(b4, board.White)

	pos := NewPosition(board.White, []Piece{whiteKing, whiteBishop})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves")
	}

	expectedMoves := moves.MoveList{}

	expectedSquares := []string{"a2", "b2", "c2", "a3", "c3", "a4", "c4"}
	for _, sqStr := range expectedSquares {
		sq := board.NewSquareOrPanic(sqStr)
		expectedMoves.AddMove(moves.NewMove(b3, sq, piece.KingType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Fatalf("Expected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Fatalf("Expected moves %v, got %v", expectedMoves, mov)
	}
}

func TestGetMovesForBishopOnEmptyBoard(t *testing.T) {
	d4 := board.NewSquareOrPanic("d4")
	whiteBishop := NewBishop(d4, board.White)

	pos := NewPosition(board.White, []Piece{whiteBishop})

	mov, err := whiteBishop.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	expectedMoves := moves.MoveList{}

	expectedSquares := []string{
		"a1", "b2", "c3", "e5",
		"f6", "g7", "h8", "a7",
		"b6", "c5", "e3", "f2",
		"g1",
	}
	for _, sqStr := range expectedSquares {
		sq := board.NewSquareOrPanic(sqStr)
		expectedMoves.AddMove(moves.NewMove(d4, sq, piece.BishopType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Errorf("Expected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	// Compare move lists
	if equal := mov.Equals(expectedMoves); !equal {
		t.Errorf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}

func TestGetMovesForBishopWhenAnotherPieceOccupiesOneOfThePossibleSquares(t *testing.T) {
	b2 := board.NewSquareOrPanic("b2")
	whiteBishop := NewBishop(b2, board.White)

	a1 := board.NewSquareOrPanic("a1")
	whitePawn := NewPawn(a1, board.White)

	pos := NewPosition(board.White, []Piece{whiteBishop, whitePawn})

	mov, err := whiteBishop.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves")
	}

	expectedMoves := moves.MoveList{}

	expectedSquares := []string{"c3", "d4", "e5", "f6", "g7", "h8", "a3", "c1"}
	for _, sqStr := range expectedSquares {
		sq := board.NewSquareOrPanic(sqStr)
		expectedMoves.AddMove(moves.NewMove(b2, sq, piece.BishopType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Fatalf("Expected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Errorf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}

func TestGetMovesForRookWithFriendlyPieceBlocking(t *testing.T) {
	f4 := board.NewSquareOrPanic("f4")
	blackRook := NewRook(f4, board.Black)

	f3 := board.NewSquareOrPanic("f3")
	blackBishop := NewBishop(f3, board.Black)

	pos := NewPosition(board.Black, []Piece{blackRook, blackBishop})

	mov, err := blackRook.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves for rook")
	}

	expectedSquares := []string{
		"f5", "f6", "f7", "f8",
		"e4", "d4", "c4", "b4",
		"a4", "g4", "h4",
	}

	expectedMoves := moves.MoveList{}

	for _, sq := range expectedSquares {
		square := board.NewSquareOrPanic(sq)
		expectedMoves.AddMove(moves.NewMove(f4, square, piece.RookType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}

func TestGetMovesForQueen(t *testing.T) {
	h8 := board.NewSquareOrPanic("h8")
	whiteQueen := NewQueen(h8, board.White)

	h1 := board.NewSquareOrPanic("h1")
	whitePawn := NewPawn(h1, board.White)

	b2 := board.NewSquareOrPanic("b2")
	whiteKing := NewKing(b2, board.White)

	pos := NewPosition(board.White, []Piece{whiteQueen, whitePawn, whiteKing})

	mov, err := whiteQueen.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves for queen")
	}

	expectedSquares := []string{
		"h7", "h6", "h5", "h4", "h3", "h2",
		"g7", "f6", "e5", "d4", "c3",
		"g8", "f8", "e8", "d8", "c8", "b8", "a8",
	}

	expectedMoves := moves.MoveList{}

	for _, sq := range expectedSquares {
		square := board.NewSquareOrPanic(sq)
		expectedMoves.AddMove(moves.NewMove(h8, square, piece.QueenType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}

func TestGetPawnMoves(t *testing.T) {
	a2 := board.NewSquareOrPanic("a2")
	whitePawn := NewPawn(a2, board.White)

	a7 := board.NewSquareOrPanic("a7")
	blackPawn := NewPawn(a7, board.Black)

	pos := NewPosition(board.White, []Piece{whitePawn, blackPawn})

	movs, err := whitePawn.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves for white pawn")
	}

	movesBlack, err := blackPawn.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves for black pawn")
	}

	movs.AddMoveList(movesBlack)

	expectedMoves := moves.MoveList{}

	expectedSquaresForWhitePawn := []string{"a3", "a4"}
	expectedSquaresForBlackPawn := []string{"a6", "a5"}

	for _, sq := range expectedSquaresForWhitePawn {
		square := board.NewSquareOrPanic(sq)
		expectedMoves.AddMove(moves.NewMove(a2, square, piece.PawnType, false))
	}

	for _, sq := range expectedSquaresForBlackPawn {
		square := board.NewSquareOrPanic(sq)
		expectedMoves.AddMove(moves.NewMove(a7, square, piece.PawnType, false))
	}

	if equal := movs.Equals(expectedMoves); !equal {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, movs)
	}
}

func TestGetKnightMoves(t *testing.T) {
	e6 := board.NewSquareOrPanic("e6")
	blackKnight := NewKnight(e6, board.Black)

	g5 := board.NewSquareOrPanic("g5")
	blackQueen := NewQueen(g5, board.Black)

	pos := NewPosition(board.Black, []Piece{blackKnight, blackQueen})

	mov, err := blackKnight.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves for knight")
	}

	expectedMoves := moves.MoveList{}

	expectedSquares := []string{"f8", "g7", "f4", "d4", "c5", "c7", "d8"}
	for _, sq := range expectedSquares {
		square := board.NewSquareOrPanic(sq)
		expectedMoves.AddMove(moves.NewMove(e6, square, piece.KnightType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}

func TestGetKnightMovesWithKnightInCorner(t *testing.T) {
	a1 := board.NewSquareOrPanic("a1")
	whiteKnight := NewKnight(a1, board.White)

	pos := NewPosition(board.White, []Piece{whiteKnight})

	mov, err := whiteKnight.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves for knight")
	}

	expectedMoves := moves.MoveList{}

	expectedSquares := []string{"b3", "c2"}
	for _, sq := range expectedSquares {
		square := board.NewSquareOrPanic(sq)
		expectedMoves.AddMove(moves.NewMove(a1, square, piece.KnightType, false))
	}

	if mov.Len() != expectedMoves.Len() {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected %d moves, got %d", expectedMoves.Len(), mov.Len())
	}

	if equal := mov.Equals(expectedMoves); !equal {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}
