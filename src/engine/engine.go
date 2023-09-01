package engine

import (
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/position"
	"math/rand"
)

type Engine struct {
	// The position the engine is currently analysing.
	Position *position.Position
	// The current search depth.
	Depth int
	// The maximum search depth.
	MaxDepth int
}

func NewEngine() *Engine {
	return &Engine{
		Position: position.NewStartingPosition(),
		Depth:    0,
		MaxDepth: 1,
	}
}

func (e *Engine) SetPosition(pos *position.Position) {
	e.Position = pos
}

func (e *Engine) FindBestMove() (moves.Move, error) {
	allPossibleMoves, err := e.Position.GetAllPossibleMoves()
	if err != nil {
		return moves.Move{}, err
	}

	// pick one at random for now
	randomNumber := rand.Intn(allPossibleMoves.Len())
	return allPossibleMoves.GetMoves()[randomNumber], nil
}
