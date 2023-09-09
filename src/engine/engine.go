package engine

import (
	"fmt"
	"math/rand"

	"github.com/samwestmoreland/chessengine/src/eval"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/samwestmoreland/chessengine/src/position"
)

type Engine struct {
	// The position the engine is currently analysing.
	pos *position.Position
	// The current search depth.
	Depth int
	// The maximum search depth.
	MaxDepth  int
	evaluator eval.Evaluator
}

func NewEngine() *Engine {
	return &Engine{
		pos:       position.NewStartingPosition(),
		Depth:     0,
		MaxDepth:  1,
		evaluator: eval.ShannonEvaluator{},
	}
}

func (e *Engine) SetPosition(pos *position.Position) {
	e.pos = pos
}

func (e *Engine) FindBestMove() (moves.Move, error) {
	allPossibleMoves, err := e.pos.GetAllPossibleMoves()
	if err != nil {
		return moves.Move{}, fmt.Errorf("failed to get all possible moves: %w", err)
	}

	/* #nosec G404 */
	randomNumber := rand.Intn(allPossibleMoves.Len())

	return allPossibleMoves.GetMoves()[randomNumber], nil
}

func (e *Engine) Evaluate() float64 {
	return e.evaluator.Evaluate(e.pos)
}
