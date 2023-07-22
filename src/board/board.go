package board

import (
	"fmt"
)

type Square struct {
	Rank int
	File string
}

func (s Square) String() string {
	return string(s)
}

func (s Square) IsValidSquare() bool {
	return s == Empty
}
