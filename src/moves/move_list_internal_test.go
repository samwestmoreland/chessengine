package moves

import (
	"testing"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/piece"
)

func TestAddMove(t *testing.T) {
	aMove := NewMove(
		board.NewSquareOrPanic("a2"),
		board.NewSquareOrPanic("a3"),
		piece.PawnType,
		false,
	)
	aMoveList := MoveList{[]Move{aMove}}

	anotherMove := NewMove(
		board.NewSquareOrPanic("b4"),
		board.NewSquareOrPanic("b5"),
		piece.PawnType,
		false,
	)
	aMoveList.AddMove(anotherMove)

	expectedMoveList := MoveList{[]Move{aMove, anotherMove}}

	if !aMoveList.Equals(expectedMoveList) {
		t.Errorf("\nExpected:\n%vGot:\n%v", expectedMoveList.String(), aMoveList.String())
	}
}

func TestAddMoves(t *testing.T) {
	aMove := NewMove(
		board.NewSquareOrPanic("a2"),
		board.NewSquareOrPanic("a3"),
		piece.PawnType,
		false,
	)

	anotherMove := NewMove(
		board.NewSquareOrPanic("b4"),
		board.NewSquareOrPanic("b5"),
		piece.PawnType,
		false,
	)
	aMoveList := MoveList{[]Move{}}
	aMoveList.AddMoves([]Move{aMove, anotherMove})

	expectedMoveList := MoveList{[]Move{aMove, anotherMove}}

	if !aMoveList.Equals(expectedMoveList) {
		t.Errorf("\nExpected:\n%vGot:\n%v", expectedMoveList.String(), aMoveList.String())
	}
}

func TestAddMoveList(t *testing.T) {
	aMove := NewMove(
		board.NewSquareOrPanic("a2"),
		board.NewSquareOrPanic("a3"),
		piece.PawnType,
		false,
	)
	aMoveList := MoveList{[]Move{aMove}}

	anotherMove := NewMove(
		board.NewSquareOrPanic("b4"),
		board.NewSquareOrPanic("b5"),
		piece.PawnType,
		false,
	)
	anotherMoveList := MoveList{[]Move{anotherMove}}

	aMoveList.AddMoveList(anotherMoveList)

	expectedMoveList := MoveList{[]Move{aMove, anotherMove}}

	if !aMoveList.Equals(expectedMoveList) {
		t.Errorf("\nExpected:\n%vGot:\n%v", expectedMoveList.String(), aMoveList.String())
	}
}

func TestMoveListLength(t *testing.T) {
	aMove := NewMove(
		board.NewSquareOrPanic("a2"),
		board.NewSquareOrPanic("a3"),
		piece.PawnType,
		false,
	)
	aMoveList := MoveList{[]Move{aMove}}

	if aMoveList.Len() != 1 {
		t.Errorf("Expected length of 1, got %v", aMoveList.Len())
	}
}

func TestMoveListEquals(t *testing.T) {
	aMove := NewMove(
		board.NewSquareOrPanic("a2"),
		board.NewSquareOrPanic("a3"),
		piece.PawnType,
		false,
	)
	aMoveList := MoveList{[]Move{aMove}}
	aMoveListCopy := MoveList{[]Move{aMove}}

	anotherMove := NewMove(
		board.NewSquareOrPanic("b4"),
		board.NewSquareOrPanic("b5"),
		piece.PawnType,
		false,
	)
	andAnotherMove := NewMove(
		board.NewSquareOrPanic("c4"),
		board.NewSquareOrPanic("c5"),
		piece.PawnType,
		false,
	)
	moveListWithOne := MoveList{[]Move{anotherMove}}
	moveListWithTwo := MoveList{[]Move{anotherMove, andAnotherMove}}

	if aMoveList.Equals(moveListWithOne) {
		t.Errorf("Expected move lists to not be equal")
	}

	if aMoveList.Equals(moveListWithTwo) {
		t.Errorf("Expected move lists to not be equal")
	}

	if !aMoveList.Equals(aMoveListCopy) {
		t.Errorf("Expected move lists to be equal")
	}
}

func TestMoveListString(t *testing.T) {
	aMove := NewMove(
		board.NewSquareOrPanic("a2"),
		board.NewSquareOrPanic("a3"),
		piece.PawnType,
		false,
	)
	aMoveList := MoveList{[]Move{aMove}}

	expectedString := "Pawn: a2 -> a3\n\n"

	if aMoveList.String() != expectedString {
		t.Errorf("Expected string \"%v\", got \"%v\"", expectedString, aMoveList.String())
	}
}

func TestGetMovesForPieceType(t *testing.T) {
	aMove := NewMove(
		board.NewSquareOrPanic("a2"),
		board.NewSquareOrPanic("a3"),
		piece.PawnType,
		false,
	)
	aMoveList := MoveList{[]Move{aMove}}

	pawnMoveList := aMoveList.GetMovesForPieceType(piece.PawnType)
	if !pawnMoveList.Equals(aMoveList) {
		t.Errorf("Expected move lists to be equal")
	}
}
