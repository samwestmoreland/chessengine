package position

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/pieces"
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
	Colour         string       // w or b
	CastlingRights string       // KQkq
	EnPassant      board.Square // e3
	HalfMoveClock  int          // the number of half moves since the last capture or pawn advance
	FullMoveNumber int          // the number of full moves, starting at 1 and incrementing after black moves
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
	enPassant, err := board.ParseSquare(parts[3])
	if err != nil {
		return nil, fmt.Errorf("Failed to parse en passant square: %w", err)
	}
	ret.EnPassant = enPassant

	ret.HalfMoveClock, err = strconv.Atoi(parts[4])
	if err != nil {
		return nil, fmt.Errorf("Failed to parse half move clock: %w", err)
	}
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

func (f FEN) GetPiece(s board.Square) (pieces.Piece, error) {
	// Given a fen, return the piece at the given square
	// The first rank is the 8th rank, so we need to reverse the ranks
	ranks := strings.Split(f.Str, "/")
	if len(ranks) != 8 {
		return nil, fmt.Errorf("FEN must have 8 ranks")
	}

	// check the square is valid
	if err := s.Valid(); err != nil {
		return nil, err
	}

	// now we know the square is valid, we can get the rank and file
	rank := s.Rank
	file := s.File

	// reverse the ranks
	rank = 7 - rank

	// get the rank
	r := ranks[rank]

	// now we need to iterate through the rank, counting the number of empty squares
	// until we get to the file we want
	emptySquares := 0
	for _, c := range r {
		if c == ' ' {
			return nil, fmt.Errorf("FEN contains a space")
		}
		if c == '/' {
			return nil, fmt.Errorf("FEN contains a /")
		}
		if c >= '1' && c <= '8' {
			emptySquares += int(c - '0')
		} else {
			emptySquares++
		}
		if emptySquares > file {
			return nil, fmt.Errorf("FEN contains a space")
		}
		if emptySquares == file {
			// we've found the square we want
			return pieces.FromChar(c, s), nil
		}
	}

	return nil, fmt.Errorf("Failed to find piece at square %s", s.String())

}

// validateCastlingRights checks that the castling rights string is valid, returning an error if not
func validateCastlingRights(s string) error {
	if len(s) > 4 {
		return fmt.Errorf("Castling rights cannot be longer than 4 characters")
	}
	validRegex := "^[K?Q?k?q?]|-"
	if _, err := regexp.MatchString(validRegex, s); err != nil {
		return fmt.Errorf("Castling rights string is invalid: %w", err)
	}

	return nil
}
