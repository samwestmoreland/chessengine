package position

import (
	"fmt"
	"strconv"
	"strings"
)

// A position is in Forsyth–Edwards notation
// From wikipedia:
// Here is the FEN for the starting position:
// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1

// And after the move 1.e4:
// rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1

// And then after 1...c5:
// rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2

// And then after 2.Nf3:
// rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2

type FEN struct {
	Str            string
	Colour         string // w or b
	CastlingRights string // KQkq
	EnPassant      string // e3
	HalfMoveClock  int    // the number of half moves since the last capture or pawn advance
	FullMoveNumber int    // the number of full moves, starting at 1 and incrementing after black moves
}

func (f FEN) String() string {
	return f.Str
}

func ParseFEN(s string) (*FEN, error) {
	// split the string at the spaces
	parts := strings.Split(s, " ")
	if len(parts) != 6 {
		return nil, fmt.Errorf("FEN must have 6 parts")
	}
	if len(s) > 100 {
		return nil, fmt.Errorf("FEN too long")
	}

	// check the first part is valid
	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return nil, fmt.Errorf("FEN must have 8 ranks")
	}

	var ret FEN
	ret.Str = s
	ret.Colour = parts[1]
	ret.CastlingRights = parts[2]
	ret.EnPassant = parts[3]

	halfMoveClock, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil, fmt.Errorf("Failed to parse half move clock: %w", err)
	}
	ret.HalfMoveClock = halfMoveClock
	ret.FullMoveNumber, err = strconv.Atoi(parts[5])
	if err != nil {
		return nil, fmt.Errorf("Failed to parse full move number: %w", err)
	}

	if ret.Colour != "w" && ret.Colour != "b" {
		return nil, fmt.Errorf("FEN colour must be w or b")
	}

	if err := validateCastlingRights(ret.CastlingRights); err != nil {
		return nil, err
	}

	return &ret, nil
}

// validateCastlingRights checks that the castling rights string is valid, returning an error if not
func validateCastlingRights(s string) error {
	if len(s) > 4 {
		return fmt.Errorf("Castling rights cannot be longer than 4 characters")
	}
	// TODO: there are more valid strings to add
	validStrings := []string{"KQkq", "KQk", "Kk", "Qq", "K", "Q", "k", "q", "-"}
	for _, valid := range validStrings {
		if s == valid {
			return nil
		}
	}
	return nil
}
