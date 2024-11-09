package engine

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
