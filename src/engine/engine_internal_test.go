package engine

import (
	"testing"
)

func TestGetAMoveInPosition(t *testing.T) {
	// this iniitializes the engine with the starting position
	e := NewEngine()

	move, err := e.FindBestMove()
	if err != nil {
		t.Error("GetAMoveInPosition() failed to return a move")
	}

	if !move.To.Valid() {
		t.Error("GetAMoveInPosition() returned invalid move")
	}
}

func TestEvaluateStartingPosition(t *testing.T) {
	e := NewEngine()

	score := e.Evaluate()
	if score != 0 {
		t.Error("EvaluateStartingPosition() failed to return 0")
	}
}
