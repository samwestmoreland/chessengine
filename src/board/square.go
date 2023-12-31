package board

import (
	"errors"
	"fmt"
	"strconv"
)

type Square struct {
	Rank, File int
}

var (
	ErrInvalidRank   = errors.New("invalid rank")
	ErrInvalidFile   = errors.New("invalid file")
	ErrInvalidSquare = errors.New("invalid square")
	ErrNilSquare     = errors.New("nil square")
)

// NewSquare takes a string representation of a square and returns a Square.
func NewSquare(sqStr string) (Square, error) {
	if len(sqStr) != 2 {
		return Square{}, fmt.Errorf("string has incorrect length: %s: %w", sqStr, ErrInvalidSquare)
	}

	// Convert the letter into an int
	file := int(sqStr[0]) - int('a') + 1

	rank, err := strconv.Atoi(string(sqStr[1]))
	if err != nil {
		return Square{}, fmt.Errorf("invalid rank: %s: %w", string(sqStr[1]), ErrInvalidRank)
	}

	ret := Square{Rank: rank, File: file}
	if !ret.Valid() {
		return Square{}, fmt.Errorf("invalid square: %s: %w", sqStr, ErrInvalidSquare)
	}

	return ret, nil
}

func NewSquareOrPanic(sqStr string) Square {
	sq, err := NewSquare(sqStr)
	if err != nil {
		panic(err)
	}

	return sq
}

// String returns the string representation of a square. So a1, b2, etc.
func (s *Square) String() string {
	file := rune('a' + s.File - 1)

	return fmt.Sprintf("%v%d", string(file), s.Rank)
}

func (s Square) Valid() bool {
	if s.Rank < 1 || s.Rank > 8 {
		return false
	}

	if s.File < 1 || s.File > 8 {
		return false
	}

	return true
}

func (s Square) IsLightSquare() (bool, error) {
	return s.Valid() && (s.Rank+s.File)%2 == 1, nil
}

func (s Square) IsDarkSquare() (bool, error) {
	isLight, err := s.IsLightSquare()
	if err != nil {
		return false, err
	}

	return !isLight, nil
}

func (s Square) IsSameSquare(other Square) bool {
	return s.Rank == other.Rank && s.File == other.File
}

func (s Square) IsSameRank(other Square) bool {
	return s.Rank == other.Rank
}

func (s Square) IsSameFile(other Square) bool {
	return s.File == other.File
}

func (s Square) IsSameDiagonal(other Square) bool {
	return s.Rank-other.Rank == other.File-s.File
}

func (s Square) Translate(direction Direction) Square {
	switch direction {
	case North:
		return Square{Rank: s.Rank + 1, File: s.File}
	case NorthEast:
		return Square{Rank: s.Rank + 1, File: s.File + 1}
	case East:
		return Square{Rank: s.Rank, File: s.File + 1}
	case SouthEast:
		return Square{Rank: s.Rank - 1, File: s.File + 1}
	case South:
		return Square{Rank: s.Rank - 1, File: s.File}
	case SouthWest:
		return Square{Rank: s.Rank - 1, File: s.File - 1}
	case West:
		return Square{Rank: s.Rank, File: s.File - 1}
	case NorthWest:
		return Square{Rank: s.Rank + 1, File: s.File - 1}
	default:
		return s
	}
}
