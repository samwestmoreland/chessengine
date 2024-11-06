package squares_test

import (
	"testing"

	sq "github.com/samwestmoreland/chessengine/src/squares"
)

func TestStringify(t *testing.T) {
	testCases := map[int]string{
		0:  "a8",
		1:  "b8",
		2:  "c8",
		3:  "d8",
		4:  "e8",
		5:  "f8",
		6:  "g8",
		7:  "h8",
		8:  "a7",
		9:  "b7",
		10: "c7",
		11: "d7",
		12: "e7",
		13: "f7",
		14: "g7",
		15: "h7",
		16: "a6",
		17: "b6",
		18: "c6",
		19: "d6",
		55: "h2",
		56: "a1",
		57: "b1",
		58: "c1",
		59: "d1",
		60: "e1",
		61: "f1",
		62: "g1",
		63: "h1",
	}

	for square, expected := range testCases {
		actual := sq.Stringify(square)
		if actual != expected {
			t.Errorf("While stringifying %d, expected %s, got %s", square, expected, actual)
		}
	}
}
