package engine

import (
	"log"
	"math/rand/v2"

	"github.com/samwestmoreland/chessengine/internal/movegen"
	"github.com/samwestmoreland/chessengine/internal/position"
)

type Engine struct {
	// The current search depth.
	Depth int
	// The maximum search depth.
	MaxDepth int
	// evaluator eval.Evaluator
}

func NewEngine() (*Engine, error) {
	return &Engine{
		Depth:    0,
		MaxDepth: 1,
		// evaluator: eval.ShannonEvaluator{},
	}, nil
}

func (e *Engine) Search(pos *position.Position) string {
	moves := movegen.GetLegalMoves(pos)

	numMoves := len(moves)

	log.Println("Number of moves:", numMoves)

	rand := rand.IntN(numMoves)

	return moves[rand].String()
}
