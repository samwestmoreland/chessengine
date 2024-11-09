package position

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

func TestNewStateFromFEN(t *testing.T) {
	tests := []struct {
		fen   string
		error bool
	}{
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			false,
		},
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkqq - 0 1",
			true,
		},
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR g KQkq - 0 1",
			true,
		},
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 2",
			true,
		},
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNB3KBNR w KQkq - 0 1",
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			_, err := NewStateFromFEN(test.fen)
			if err != nil && !test.error {
				t.Error(err)
			}
		})
	}
}
