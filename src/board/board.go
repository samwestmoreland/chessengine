package board

import (
	"fmt"

	"github.com/uber-go/zap"
)

type Square struct {
	Rank, File int
}

func (s Square) String() string {
	return fmt.Sprintf("%s%d", s.File, s.Rank)
}

func (s Square) IsValidSquare() bool {
	return s.Rank >= 1 && s.Rank <= 8 && s.File >= 'a' && s.File <= 'h'
}

func (s Square) IsLightSquare() bool {
	return (s.Rank+s.File)%2 == 0
}

func (s Square) IsDarkSquare() bool {
	return !s.IsLightSquare()
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
	rank := int(s[1])
	file := int(s[0] - 'a' + 1)
	return Square{rank, file}, nil
}
