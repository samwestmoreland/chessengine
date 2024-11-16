package move

import (
	"fmt"

	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

type Move struct {
	From           int
	To             int
	PromotionPiece string
}

func (m Move) String() string {
	ret := fmt.Sprintf("%s%s", sq.Stringify(m.From), sq.Stringify(m.To))

	if m.PromotionPiece != "" {
		ret += fmt.Sprintf("%s", m.PromotionPiece)
	}

	return ret
}
