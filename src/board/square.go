package board

import (
	"fmt"
)

type Square struct {
	Rank, File int
}

func (s Square) String() string {
	file := string('a' + s.File - 1)
	return fmt.Sprintf("%s%d", file, s.Rank)
}

func (s Square) CheckValidity() error {
	if s.Rank < 1 || s.Rank > 8 {
		return fmt.Errorf("Invalid rank: %d", s.Rank)
	}
	if s.File < 1 || s.File > 8 {
		return fmt.Errorf("Invalid file: %d", s.File)
	}
	return nil
}

func (s Square) IsLightSquare() (bool, error) {
	if err := s.CheckValidity(); err != nil {
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
	return square, square.CheckValidity()
}