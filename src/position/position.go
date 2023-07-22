package position

import (
	"fmt"
)

// A position is in Forsyth–Edwards notation
// Here is the FEN for the starting position:
//
// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
// And after the move 1.e4:
//
// rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1
// And then after 1...c5:
//
// rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2
// And then after 2.Nf3:
//
// rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2

type FEN struct {
	Raw string
}

func (f FEN) String() string {
	return f.Raw
}

func (f FEN) IsValid() bool {
	// things we need to check:
	// 1. 8 rows
	return true
}

func GetFEN(s string) (*FEN, error) {
	var ret FEN
	if len(s) == 0 {
		return nil, fmt.Errorf("Empty string")
	}
	if len(s) > 100 {
		return nil, fmt.Errorf("String too long")
	}
	ret.Raw = s
	return &ret, nil
}
