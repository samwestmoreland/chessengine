package position

import (
	"strings"
	"testing"
)

func TestParseCastlingRights(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			result, err := parseCastlingRights(test.input)
			if err != nil {
				t.Error(err)
			}

			if result != test.output {
				t.Errorf("Expected %d, got %d", test.output, result)
			}
		})
	}
}

func TestNewPositionFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		fen           string
		expectPanic   bool
		expectError   bool
		errorContains string
	}{
		{
			"starting position",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			false,
			false,
			"",
		},
		{
			"invalid castling rights",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkqq - 0 1",
			false,
			true,
			"expected castling rights",
		},
		{
			"invalid side",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR g KQkq - 0 1",
			false,
			true,
			"invalid side",
		},
		{
			"extra fields",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 2",
			false,
			true,
			"FEN must have",
		},
		{
			"too many squares",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNB3KBNR w KQkq - 0 1",
			true,
			false,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			pos, err := NewPositionFromFEN(tt.fen)

			if tt.expectPanic {
				t.Error("expected panic but got none")
			}

			if err != nil {
				if !tt.expectError {
					t.Errorf("unexpected error: %v", err)
				} else if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("error %q does not contain %q", err.Error(), tt.errorContains)
				}
			} else if tt.expectError {
				t.Error("expected error but got none")
			}

			if pos != nil && tt.expectError {
				t.Error("got position when expecting error")
			}
		})
	}
}
