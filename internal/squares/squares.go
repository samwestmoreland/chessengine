package squares

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	A8 = iota
	B8
	C8
	D8
	E8
	F8
	G8
	H8

	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7

	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6

	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5

	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4

	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3

	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2

	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1

	NoSquare
)

type Square uint8

func Stringify(square Square) string {
	if square > 63 {
		return "-"
	}

	var ret strings.Builder

	rank := 8 - square/8

	file := square%8 + 1

	ret.WriteString(fmt.Sprintf("%c", 'a'+file-1))
	ret.WriteString(strconv.Itoa(int(rank)))

	return ret.String()
}

func ParseString(square string) (Square, error) {
	if len(square) != 2 {
		return 0, fmt.Errorf("invalid square format: %s", square)
	}

	rank, err := strconv.Atoi(square[1:2])
	if err != nil {
		return 0, fmt.Errorf("invalid rank: %s", square[1:2])
	}

	if rank < 1 || rank > 8 {
		return 0, fmt.Errorf("rank out of range: %d", rank)
	}

	rankIndex := 8 - rank

	file := strings.ToLower(square[0:1])
	if file < "a" || file > "h" {
		return 0, fmt.Errorf("invalid file: %s", file)
	}

	fileIndex := file[0] - 'a'

	index := rankIndex*8 + int(fileIndex)

	return Square(index), nil
}

func OnBoard(square Square) bool {
	return square <= 63
}
