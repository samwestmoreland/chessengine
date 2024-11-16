package move

import (
	"fmt"

	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

// Moves are represented as bitmasks:
//
//	0000 0000 0000 0000 0011 1111    source square
//	0000 0000 0000 1111 1100 0000    target square
//	0000 0000 1111 0000 0000 0000    piece
//	0000 1111 0000 0000 0000 0000    promotion piece
//	0001 0000 0000 0000 0000 0000    capture flag
//	0010 0000 0000 0000 0000 0000    double push flag
//	0100 0000 0000 0000 0000 0000    en passant flag
//	1000 0000 0000 0000 0000 0000    castling flag
type Move uint32

func (m Move) String() string {
	ret := fmt.Sprintf("%s%s", sq.Stringify(m.From), sq.Stringify(m.To))

	if m.PromotionPiece != "" {
		ret += fmt.Sprintf("%s", m.PromotionPiece)
	}

	return ret
}
