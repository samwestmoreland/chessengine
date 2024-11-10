package squares_test

import (
	"testing"

	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

var tests = []struct {
	integerRepresentation int
	stringRepresentation  string
}{
	{sq.A1, "a1"},
	{sq.A2, "a2"},
	{sq.A3, "a3"},
	{sq.A4, "a4"},
	{sq.A5, "a5"},
	{sq.A6, "a6"},
	{sq.A7, "a7"},
	{sq.A8, "a8"},
	{sq.B1, "b1"},
	{sq.B2, "b2"},
	{sq.B3, "b3"},
	{sq.B4, "b4"},
	{sq.B5, "b5"},
	{sq.B6, "b6"},
	{sq.B7, "b7"},
	{sq.B8, "b8"},
	{sq.C1, "c1"},
	{sq.C2, "c2"},
	{sq.C3, "c3"},
	{sq.C4, "c4"},
	{sq.C5, "c5"},
	{sq.C6, "c6"},
	{sq.C7, "c7"},
	{sq.C8, "c8"},
	{sq.D1, "d1"},
	{sq.D2, "d2"},
	{sq.D3, "d3"},
	{sq.D4, "d4"},
	{sq.D5, "d5"},
	{sq.D6, "d6"},
	{sq.D7, "d7"},
	{sq.D8, "d8"},
	{sq.E1, "e1"},
	{sq.E2, "e2"},
	{sq.E3, "e3"},
	{sq.E4, "e4"},
	{sq.E5, "e5"},
	{sq.E6, "e6"},
	{sq.E7, "e7"},
	{sq.E8, "e8"},
	{sq.F1, "f1"},
	{sq.F2, "f2"},
	{sq.F3, "f3"},
	{sq.F4, "f4"},
	{sq.F5, "f5"},
	{sq.F6, "f6"},
	{sq.F7, "f7"},
	{sq.F8, "f8"},
	{sq.G1, "g1"},
	{sq.G2, "g2"},
	{sq.G3, "g3"},
	{sq.G4, "g4"},
	{sq.G5, "g5"},
	{sq.G6, "g6"},
	{sq.G7, "g7"},
	{sq.G8, "g8"},
	{sq.H1, "h1"},
	{sq.H2, "h2"},
	{sq.H3, "h3"},
	{sq.H4, "h4"},
	{sq.H5, "h5"},
	{sq.H6, "h6"},
	{sq.H7, "h7"},
	{sq.H8, "h8"},
}

func TestStringify(t *testing.T) {
	for _, test := range tests {
		t.Run(test.stringRepresentation, func(t *testing.T) {
			result := sq.Stringify(test.integerRepresentation)
			if result != test.stringRepresentation {
				t.Errorf("Expected %s, got %s", test.stringRepresentation, result)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	for _, test := range tests {
		t.Run(test.stringRepresentation, func(t *testing.T) {
			result, err := sq.ToInt(test.stringRepresentation)
			if err != nil {
				t.Error(err)
			}

			if result != test.integerRepresentation {
				t.Errorf("Expected %d, got %d", test.integerRepresentation, result)
			}
		})
	}
}

func TestOnBoard(t *testing.T) {
	testCases := []struct {
		square int
		want   bool
	}{
		{sq.A1, true},
		{sq.C3, true},
		{75, false},
		{-1, false},
	}

	for _, tc := range testCases {
		got := sq.OnBoard(tc.square)
		if got != tc.want {
			t.Errorf("OnBoard(%d) = %v, want %v", tc.square, got, tc.want)
		}
	}
}
