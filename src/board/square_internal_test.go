package board

import (
	"testing"
)

func TestIsValidSquare(t *testing.T) {
	squaresToTest := map[Square]bool{
		{File: -1, Rank: 0}:   false,
		{File: 0, Rank: -1}:   false,
		{File: 0, Rank: 0}:    false,
		{File: 1, Rank: 1}:    true,
		{File: 7, Rank: 7}:    true,
		{File: 8, Rank: 7}:    true,
		{File: 7, Rank: 8}:    true,
		{File: 8, Rank: 8}:    true,
		{File: 10, Rank: 8}:   false,
		{File: 10, Rank: 100}: false,
	}

	for square, expected := range squaresToTest {
		actual := square.Valid()
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestSquareColour(t *testing.T) {
	darkSquares := []Square{
		{File: 1, Rank: 1},
		{File: 2, Rank: 2},
		{File: 3, Rank: 3},
		{File: 4, Rank: 4},
		{File: 5, Rank: 5},
		{File: 6, Rank: 6},
		{File: 7, Rank: 7},
		{File: 8, Rank: 8},
		{File: 1, Rank: 3},
		{File: 2, Rank: 4},
		{File: 3, Rank: 5},
		{File: 4, Rank: 6},
		{File: 5, Rank: 7},
		{File: 6, Rank: 8},
	}
	for _, square := range darkSquares {
		if dark, err := square.IsDarkSquare(); err != nil || !dark {
			t.Errorf("Expected %v to be dark", square)
		}
	}

	lightSquares := []Square{
		{File: 1, Rank: 2},
		{File: 2, Rank: 1},
		{File: 3, Rank: 2},
		{File: 4, Rank: 3},
		{File: 5, Rank: 4},
		{File: 6, Rank: 5},
		{File: 7, Rank: 6},
		{File: 8, Rank: 7},
		{File: 2, Rank: 3},
		{File: 3, Rank: 4},
		{File: 4, Rank: 5},
		{File: 5, Rank: 6},
		{File: 6, Rank: 7},
		{File: 7, Rank: 8},
	}
	for _, square := range lightSquares {
		if light, err := square.IsLightSquare(); err != nil || !light {
			t.Errorf("Expected %v to be light", square)
		}
	}
}
