package position

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/samwestmoreland/chessengine/src/board"
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

var ErrInvalidFEN = errors.New("invalid FEN")

// FEN is a struct representing a position in Forsyth–Edwards notation.
type FEN struct {
	Str            string
	Colour         string       // w or b
	CastlingRights string       // KQkq
	EnPassant      board.Square // e3
	HalfMoveClock  int          // the number of half moves since the last capture or pawn advance
	FullMoveNumber int          // the number of full moves, starting at 1 and incrementing after black moves
}

// String returns the FEN as a string.
func (f FEN) String() string {
	return f.Str
}

// ParseFEN parses a FEN string, returning a FEN struct.
func ParseFEN(fenstr string) (*FEN, error) {
	// split the string at the spaces
	parts := strings.Split(fenstr, " ")
	if len(parts) != 6 {
		return nil, fmt.Errorf("FEN must have 6 parts: %w", ErrInvalidFEN)
	}

	if len(fenstr) > 100 {
		return nil, fmt.Errorf("FEN too long: %w", ErrInvalidFEN)
	}

	// check the first part is valid
	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return nil, fmt.Errorf("FEN must have 8 ranks: %w", ErrInvalidFEN)
	}

	var ret FEN
	ret.Str = fenstr
	ret.Colour = parts[1]
	ret.CastlingRights = parts[2]

	var enPassant = &board.Square{}

	if parts[3] != "-" {
		var err error

		enPassant, err = board.NewSquare(parts[3])
		if err != nil {
			return nil, fmt.Errorf("failed to parse en passant square: %w", err)
		}
	}

	ret.EnPassant = *enPassant

	var err error

	ret.HalfMoveClock, err = strconv.Atoi(parts[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse half move clock: %w", err)
	}

	ret.FullMoveNumber, err = strconv.Atoi(parts[5])
	if err != nil {
		return nil, fmt.Errorf("failed to parse full move number: %w", err)
	}

	if ret.Colour != "w" && ret.Colour != "b" {
		return nil, fmt.Errorf("FEN colour must be w or b, got %s: %w", ret.Colour, ErrInvalidFEN)
	}

	if err := validateCastlingRights(ret.CastlingRights); err != nil {
		return nil, err
	}

	return &ret, nil
}

// GetPiece returns the piece at the given square, given a FEN.
func (f FEN) GetPiece(square board.Square) (Piece, error) {
	// Given a fen, return the piece at the given square
	// The first rank is the 8th rank, so we need to reverse the ranks
	ranks := strings.Split(f.Str, "/")
	if len(ranks) != 8 {
		return nil, fmt.Errorf("FEN must have 8 ranks: %w", ErrInvalidFEN)
	}

	// check the square is valid
	if err := square.Valid(); err != nil {
		return nil, fmt.Errorf("got invalid square while getting piece: %w", err)
	}

	// now we know the square is valid, we can get the rank and file
	rank := square.Rank
	file := square.File

	// reverse the ranks
	rank = 8 - rank

	// get the rank
	rankStr := ranks[rank]

	// now we need to iterate through the rank, counting the number of empty squares
	// until we get to the file we want
	emptySquares := 0

	for _, character := range rankStr {
		if character == ' ' {
			return nil, fmt.Errorf("FEN contains a space: %w", ErrInvalidFEN)
		}

		if character == '/' {
			return nil, fmt.Errorf("FEN contains a /: %w", ErrInvalidFEN)
		}

		if character >= '1' && character <= '8' {
			emptySquares += int(character - '0')
		} else {
			emptySquares++
		}

		if emptySquares > file {
			return nil, fmt.Errorf("FEN contains a space: %w", ErrInvalidFEN)
		}

		if emptySquares == file {
			// we've found the square we want
			return FromChar(character, &square), nil
		}
	}

	return nil, fmt.Errorf("Failed to find piece at square %s: %w", square.String(), ErrInvalidFEN)
}

// validateCastlingRights checks that the castling rights string is valid, returning an error if not.
func validateCastlingRights(castlingStr string) error {
	if len(castlingStr) > 4 {
		return fmt.Errorf("castling rights cannot be longer than 4 characters: %w", ErrInvalidFEN)
	}

	validRegex := "^[K?Q?k?q?]|-"
	if _, err := regexp.MatchString(validRegex, castlingStr); err != nil {
		return fmt.Errorf("castling rights string is invalid: %w", err)
	}

	return nil
}
