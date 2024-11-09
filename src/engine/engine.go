package engine

import (
	"fmt"

	"github.com/samwestmoreland/chessengine/src/bitboard"
	"github.com/samwestmoreland/chessengine/src/eval"
)

type Engine struct {
	// The state of the current position.
	state *bitboard.State
	// The current search depth.
	Depth int
	// The maximum search depth.
	MaxDepth  int
	evaluator eval.Evaluator
}

func NewEngine() (*Engine, error) {
	state, err := bitboard.NewState()
	if err != nil {
		return nil, fmt.Errorf("failed to create new state: %w", err)
	}

	return &Engine{
		state:     state,
		Depth:     0,
		MaxDepth:  1,
		evaluator: eval.ShannonEvaluator{},
	}, nil
}

func (e *Engine) SetState(state *bitboard.State) {
	e.state = state
}

// func (e *Engine) SetPosition(pos *position.Position) {
// 	e.pos = pos
// }

// func (e *Engine) FindBestMove() (moves.Move, error) {
// 	allPossibleMoves, err := e.pos.GetAllPossibleMoves()
// 	if err != nil {
// 		return moves.Move{}, fmt.Errorf("failed to get all possible moves: %w", err)
// 	}
//
// 	/* #nosec G404 */
// 	randomNumber := rand.Intn(allPossibleMoves.Len())
//
// 	return allPossibleMoves.GetMoves()[randomNumber], nil
// }
//
// func (e *Engine) Evaluate() float64 {
// 	return e.evaluator.Evaluate(e.pos)
// }
