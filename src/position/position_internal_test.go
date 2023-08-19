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

func TestGetMovesForKingOnEmptyBoard(t *testing.T) {
	sqStr := "e4"

	square, err := board.NewSquare(sqStr)
	if err != nil {
		t.Fatalf("Failed to create square %s: %v", sqStr, err)
	}

	whiteKing := NewKing(square, board.White)
	pos := NewPosition(board.White, []Piece{whiteKing})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	d5, _ := board.NewSquare("d5")
	e5, _ := board.NewSquare("e5")
	f5, _ := board.NewSquare("f5")
	d4, _ := board.NewSquare("d4")
	f4, _ := board.NewSquare("f4")
	d3, _ := board.NewSquare("d3")
	e3, _ := board.NewSquare("e3")
	f3, _ := board.NewSquare("f3")

	expectedMoves := []moves.Move{
		{From: square, To: d5, PieceType: piece.KingType},
		{From: square, To: e5, PieceType: piece.KingType},
		{From: square, To: f5, PieceType: piece.KingType},
		{From: square, To: d4, PieceType: piece.KingType},
		{From: square, To: f4, PieceType: piece.KingType},
		{From: square, To: d3, PieceType: piece.KingType},
		{From: square, To: e3, PieceType: piece.KingType},
		{From: square, To: f3, PieceType: piece.KingType},
	}

	if len(mov) != len(expectedMoves) {
		t.Errorf("expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	equal := moves.MoveListsEqual(mov, expectedMoves)
	if !equal {
		t.Fatalf("expected moves %v, got %v", expectedMoves, mov)
	}
}

func TestGetMovesForKingOnEmptyBoardInCorner(t *testing.T) {
	sqStr := "a1"

	square, err := board.NewSquare(sqStr)
	if err != nil {
		t.Fatalf("Failed to create square %s: %v", sqStr, err)
	}

	whiteKing := NewKing(square, board.White)

	pos := NewPosition(board.White, []Piece{whiteKing})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	b1, _ := board.NewSquare("b1")
	a2, _ := board.NewSquare("a2")
	b2, _ := board.NewSquare("b2")

	expectedMoves := []moves.Move{
		{From: square, To: b1, PieceType: piece.KingType},
		{From: square, To: a2, PieceType: piece.KingType},
		{From: square, To: b2, PieceType: piece.KingType},
	}

	if len(mov) != len(expectedMoves) {
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	equal := moves.MoveListsEqual(mov, expectedMoves)
	if !equal {
		t.Fatalf("Expected moves %v, got %v", expectedMoves, mov)
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

func TestGetMovesForKingWhenAnotherPieceOccupiesOneOfThePossibleSquares(t *testing.T) {
	b3, _ := board.NewSquare("b3")
	b4, _ := board.NewSquare("b4")
	whiteKing := NewKing(b3, board.White)
	whiteBishop := NewBishop(b4, board.White)
	pos := NewPosition(board.White, []Piece{whiteKing, whiteBishop})

	mov, err := whiteKing.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves")
	}

	a2, _ := board.NewSquare("a2")
	b2, _ := board.NewSquare("b2")
	c2, _ := board.NewSquare("c2")
	a3, _ := board.NewSquare("a3")
	c3, _ := board.NewSquare("c3")
	a4, _ := board.NewSquare("a4")
	c4, _ := board.NewSquare("c4")

	expectedMoves := []moves.Move{
		{From: b3, To: a2, PieceType: piece.KingType},
		{From: b3, To: b2, PieceType: piece.KingType},
		{From: b3, To: c2, PieceType: piece.KingType},
		{From: b3, To: a3, PieceType: piece.KingType},
		{From: b3, To: c3, PieceType: piece.KingType},
		{From: b3, To: a4, PieceType: piece.KingType},
		{From: b3, To: c4, PieceType: piece.KingType},
	}

	if len(mov) != len(expectedMoves) {
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	equal := moves.MoveListsEqual(mov, expectedMoves)
	if !equal {
		t.Fatalf("Expected moves %v, got %v", expectedMoves, mov)
	}
}

func TestGetMovesForBishopOnEmptyBoard(t *testing.T) {
	d4, _ := board.NewSquare("d4")
	whiteBishop := NewBishop(d4, board.White)
	pos := NewPosition(board.White, []Piece{whiteBishop})

	mov, err := whiteBishop.GetMoves(pos)
	if err != nil {
		t.Errorf("Error while getting moves")
	}

	a1, _ := board.NewSquare("a1")
	b2, _ := board.NewSquare("b2")
	c3, _ := board.NewSquare("c3")
	e5, _ := board.NewSquare("e5")
	f6, _ := board.NewSquare("f6")
	g7, _ := board.NewSquare("g7")
	h8, _ := board.NewSquare("h8")
	a7, _ := board.NewSquare("a7")
	b6, _ := board.NewSquare("b6")
	c5, _ := board.NewSquare("c5")
	e3, _ := board.NewSquare("e3")
	f2, _ := board.NewSquare("f2")
	g1, _ := board.NewSquare("g1")

	expectedMoves := []moves.Move{
		{From: d4, To: a1, PieceType: piece.BishopType},
		{From: d4, To: b2, PieceType: piece.BishopType},
		{From: d4, To: c3, PieceType: piece.BishopType},
		{From: d4, To: e5, PieceType: piece.BishopType},
		{From: d4, To: f6, PieceType: piece.BishopType},
		{From: d4, To: g7, PieceType: piece.BishopType},
		{From: d4, To: h8, PieceType: piece.BishopType},
		{From: d4, To: a7, PieceType: piece.BishopType},
		{From: d4, To: b6, PieceType: piece.BishopType},
		{From: d4, To: c5, PieceType: piece.BishopType},
		{From: d4, To: e3, PieceType: piece.BishopType},
		{From: d4, To: f2, PieceType: piece.BishopType},
		{From: d4, To: g1, PieceType: piece.BishopType},
	}

	if len(mov) != len(expectedMoves) {
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	// Compare move lists
	equal := moves.MoveListsEqual(mov, expectedMoves)
	if !equal {
		t.Errorf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}

func TestGetMovesForBishopWhenAnotherPieceOccupiesOneOfThePossibleSquares(t *testing.T) {
	b2, _ := board.NewSquare("b2")
	a1, _ := board.NewSquare("a1")
	whiteBishop := NewBishop(b2, board.White)
	whitePawn := NewPawn(a1, board.White)
	pos := NewPosition(board.White, []Piece{whiteBishop, whitePawn})

	mov, err := whiteBishop.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves")
	}

	c3, _ := board.NewSquare("c3")
	d4, _ := board.NewSquare("d4")
	e5, _ := board.NewSquare("e5")
	f6, _ := board.NewSquare("f6")
	g7, _ := board.NewSquare("g7")
	h8, _ := board.NewSquare("h8")
	c1, _ := board.NewSquare("c1")
	a3, _ := board.NewSquare("a3")

	expectedMoves := []moves.Move{
		{From: b2, To: c3, PieceType: piece.BishopType},
		{From: b2, To: d4, PieceType: piece.BishopType},
		{From: b2, To: e5, PieceType: piece.BishopType},
		{From: b2, To: f6, PieceType: piece.BishopType},
		{From: b2, To: g7, PieceType: piece.BishopType},
		{From: b2, To: h8, PieceType: piece.BishopType},
		{From: b2, To: c1, PieceType: piece.BishopType},
		{From: b2, To: a3, PieceType: piece.BishopType},
	}

	if len(mov) != len(expectedMoves) {
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	equal := moves.MoveListsEqual(mov, expectedMoves)
	if !equal {
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

	expectedMoves := make([]moves.Move, 0, len(expectedSquares))

	for _, sq := range expectedSquares {
		square, _ := board.NewSquare(sq)
		expectedMoves = append(expectedMoves, moves.Move{From: f4, To: square, PieceType: piece.RookType})
	}

	if len(mov) != len(expectedMoves) {
		t.Logf("\n%v", pos.String())
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	if equal := moves.MoveListsEqual(mov, expectedMoves); !equal {
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

	expectedMoves := make([]moves.Move, 0, len(expectedSquares))

	for _, sq := range expectedSquares {
		square := board.NewSquareOrPanic(sq)
		expectedMoves = append(expectedMoves, moves.Move{From: h8, To: square, PieceType: piece.QueenType})
	}

	if len(mov) != len(expectedMoves) {
		t.Logf("\n%v", pos.String())
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	if equal := moves.MoveListsEqual(mov, expectedMoves); !equal {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
	}
}

func TestGetPawnMoves(t *testing.T) {
	a2 := board.NewSquareOrPanic("a2")
	pawn1 := NewPawn(a2, board.White)

	pos := NewPosition(board.White, []Piece{pawn1})

	mov, err := pawn1.GetMoves(pos)
	if err != nil {
		t.Fatalf("Error while getting moves for pawn")
	}

	expectedSquares := []string{
		"a3", "a4",
	}

	expectedMoves := make([]moves.Move, 0, len(expectedSquares))

	for _, sq := range expectedSquares {
		square := board.NewSquareOrPanic(sq)
		expectedMoves = append(expectedMoves, moves.Move{From: a2, To: square, PieceType: piece.PawnType})
	}

	if len(mov) != len(expectedMoves) {
		t.Logf("\n%v", pos.String())
		t.Fatalf("Expected %d moves, got %d", len(expectedMoves), len(mov))
	}

	if equal := moves.MoveListsEqual(mov, expectedMoves); !equal {
		t.Logf("\n%v", pos.String())
		t.Fatalf("\nExpected moves:\n%v\nGot:\n%v", expectedMoves, mov)
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
