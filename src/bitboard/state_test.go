package bitboard

import (
	"testing"
)

func TestParseCastlingRights(t *testing.T) {
	tests := []struct {
		input  string
		output uint8
	}{
		{"K", 8},
		{"Q", 4},
		{"k", 2},
		{"q", 1},
		{"kq", 3},
		{"KQ", 12},
		{"", 0},
		{"KQk", 14},
		{"KQkq", 15},
		{"-", 0},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := parseCastlingRights(test.input)
			if result != test.output {
				t.Errorf("Expected %d, got %d", test.output, result)
			}
		})
	}
}
