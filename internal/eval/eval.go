package eval

import (
	"github.com/samwestmoreland/chessengine/internal/position"
)

// Given a position, evaluate the position and return a score. The score is
// positive if the position is good for white, negative if it is good for black.

type Evaluator interface {
	Evaluate(pos *position.Position) float64
}
