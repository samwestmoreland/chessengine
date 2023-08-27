package engine

import (
	"github.com/samwestmoreland/chessengine/src/position"
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
	return &Engine{}
}

func (e *Engine) SetPosition(pos *position.Position) {
	e.Position = pos
}
