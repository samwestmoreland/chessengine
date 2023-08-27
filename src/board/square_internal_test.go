package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestNewSquare(t *testing.T) {
	invalidStrings := []string{
		"a1x", "", "x", "p4",
		"a9", "i9", "$1", "a$",
		"!!", "H1", "44", "a",
	}

	assert := assert.New(t)

	for _, str := range invalidStrings {
		if _, err := NewSquare(str); err == nil {
			t.Errorf("Expected error for %s", str)
		}

		assert.Panics(func() {
			NewSquareOrPanic(str)
		})
	}

	validStrings := []string{
		"a1", "a2", "a3", "a4",
		"a5", "a6", "a7", "a8",
	}

	for _, str := range validStrings {
		if _, err := NewSquare(str); err != nil {
			t.Errorf("Expected no error for %s", str)
		}

		assert.NotPanics(func() {
			NewSquareOrPanic(str)
		})
	}
}

func TestIsSameSquare(t *testing.T) {
	squaresToTest := map[Square]map[Square]bool{
		{File: 1, Rank: 1}: {
			{File: 1, Rank: 1}: true,
			{File: 1, Rank: 2}: false,
			{File: 2, Rank: 1}: false,
			{File: 2, Rank: 2}: false,
		},
		{File: 1, Rank: 2}: {
			{File: 1, Rank: 1}: false,
			{File: 1, Rank: 2}: true,
			{File: 2, Rank: 1}: false,
			{File: 2, Rank: 2}: false,
		},
		{File: 2, Rank: 1}: {
			{File: 1, Rank: 1}: false,
			{File: 1, Rank: 2}: false,
			{File: 2, Rank: 1}: true,
			{File: 2, Rank: 2}: false,
		},
		{File: 2, Rank: 2}: {
			{File: 1, Rank: 1}: false,
			{File: 1, Rank: 2}: false,
			{File: 2, Rank: 1}: false,
			{File: 2, Rank: 2}: true,
		},
	}

	for square, squares := range squaresToTest {
		for otherSquare, expected := range squares {
			actual := square.IsSameSquare(otherSquare)
			if actual != expected {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		}
	}
}
