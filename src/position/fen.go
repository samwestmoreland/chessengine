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
	Position       string
	Colour         board.Colour // w or b
	CastlingRights string       // KQkq
	EnPassant      board.Square // e3
	HalfMoveClock  int          // the number of half moves since the last capture or pawn advance
	FullMoveNumber int          // the number of full moves, starting at 1 and incrementing after black moves
}

// NewFEN returns a new FEN struct with the starting position.
func NewFEN() *FEN {
	return &FEN{
		Str:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w",
		Position:       "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
		Colour:         board.White,
		CastlingRights: "KQkq",
		EnPassant:      board.Square{File: 0, Rank: 0},
		HalfMoveClock:  0,
		FullMoveNumber: 1,
	}
}

// String returns the FEN as a string.
func (f FEN) String() string {
	return f.Str
}

func (f FEN) GetTurn() board.Colour {
	return f.Colour
}

func splitFEN(fenstr string) ([]string, error) {
	parts := strings.Split(fenstr, " ")
	if len(parts) != 6 {
		return nil, fmt.Errorf("FEN must have 6 parts: %w", ErrInvalidFEN)
	}

	return parts, nil
}

// ParseFEN parses a FEN string, returning a FEN struct.
func ParseFEN(fenstr string) (*FEN, error) {
	parts, err := splitFEN(fenstr)
	if err != nil {
		return nil, err
	}

	var ret FEN

	position := parts[0]
	if err := validatePosition(position); err != nil {
		return nil, err
	}

	ret.Str = fenstr
	ret.Colour = board.ColourFromString(parts[1])

	castlingRights := parts[2]
	if err := validateCastlingRights(castlingRights); err != nil {
		return nil, err
	}

	ret.CastlingRights = castlingRights
	enPassant := board.Square{File: 0, Rank: 0}

	if parts[3] != "-" {
		var err error

		enPassant, err = board.NewSquare(parts[3])
		if err != nil {
			return nil, fmt.Errorf("failed to parse en passant square: %w", err)
		}
	}

	ret.EnPassant = enPassant

	ret.HalfMoveClock, err = strconv.Atoi(parts[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse half move clock: %w", err)
	}

	ret.FullMoveNumber, err = strconv.Atoi(parts[5])
	if err != nil {
		return nil, fmt.Errorf("failed to parse full move number: %w", err)
	}

	if ret.Colour != board.White && ret.Colour != board.Black {
		return nil, fmt.Errorf("FEN colour must be w or b, got %s: %w", ret.Colour, ErrInvalidFEN)
	}

	return &ret, nil
}

// validatePosition checks that the position part of the FEN is valid.
func validatePosition(position string) error {
	// check the first part is valid
	ranks := strings.Split(position, "/")
	if len(ranks) != 8 {
		return fmt.Errorf("FEN must have 8 ranks: %w", ErrInvalidFEN)
	}

	for _, rank := range ranks {
		// check the rank is valid
		if err := validateRank(rank); err != nil {
			return err
		}
	}

	return nil
}

func validateRank(rank string) error {
	numFiles := 0

	for _, character := range rank {
		if character == ' ' {
			return fmt.Errorf("FEN contains a space: %w", ErrInvalidFEN)
		}

		if character == '/' {
			return fmt.Errorf("FEN contains a /: %w", ErrInvalidFEN)
		}

		if character >= '1' && character <= '8' {
			numFiles += int(character - '0')
		} else {
			numFiles++
		}
	}

	if numFiles != 8 {
		return fmt.Errorf("FEN must have 8 files: %w", ErrInvalidFEN)
	}

	return nil
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
	if !square.Valid() {
		return nil, fmt.Errorf("got invalid square while getting piece: %w", board.ErrInvalidSquare)
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
			return FromChar(character, square), nil
		}
	}

	return nil, fmt.Errorf("failed to find piece at square %s: %w", square.String(), ErrInvalidFEN)
}

// validateCastlingRights checks that the castling rights string is valid, returning an error if not.
func validateCastlingRights(castlingStr string) error {
	if len(castlingStr) > 4 {
		return fmt.Errorf("castling rights cannot be longer than 4 characters: %w", ErrInvalidFEN)
	}

	reg := regexp.MustCompile("^[KQkq]+|-")

	match := reg.FindString(castlingStr)
	if match != castlingStr {
		return fmt.Errorf("castling rights string is invalid: %w", ErrInvalidFEN)
	}

	return nil
}
