package board

import (
	"fmt"
	"strconv"
)

type Square struct {
	Rank, File int
}

// newSquare takes a string representation of a square and returns a Square
func NewSquare(sqStr string) (*Square, error) {
	if len(sqStr) != 2 {
		return nil, fmt.Errorf("newSquare expects a string of length 2")
	}

	// Convert the letter into an int
	file := int(sqStr[0]) - int('a') + 1

	rank, err := strconv.Atoi(string(sqStr[1]))
	if err != nil {
		return nil, fmt.Errorf("Invalid rank %v: %w", sqStr[1], err)
	}

	ret := &Square{Rank: rank, File: file}
	if err = ret.Valid(); err != nil {
		return nil, fmt.Errorf("Invalid square: %w", err)
	}

	return ret, nil
}

func (s *Square) String() string {
	file := string('a' + s.File - 1)
	return fmt.Sprintf("%s%d", file, s.Rank)
}

func (s *Square) Valid() error {
	if s == nil {
		return fmt.Errorf("Square is nil")
	}
	if s.Rank < 1 || s.Rank > 8 {
		return fmt.Errorf("Invalid rank: %d", s.Rank)
	}
	if s.File < 1 || s.File > 8 {
		return fmt.Errorf("Invalid file: %d", s.File)
	}
	return nil
}

func (s Square) IsLightSquare() (bool, error) {
	if err := s.Valid(); err != nil {
		return false, err
	}

	return (s.Rank+s.File)%2 == 1, nil
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
	return s.Rank-other.Rank == int(other.File)-int(s.File)
}

func ParseSquare(s string) (Square, error) {
	if len(s) != 2 {
		if s == "-" {
			return Square{}, nil
		}
		return Square{}, fmt.Errorf("Invalid square: %s", s)
	}
	rank := int(s[1] - '0')
	file := int(s[0] - 'a' + 1)
	square := Square{rank, file}
	fmt.Printf("Parsed square: %s\n", square)
	return square, square.Valid()
}
