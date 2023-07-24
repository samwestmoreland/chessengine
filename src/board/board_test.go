package board

import (
	"fmt"
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
		var actual bool
		var err error
		err = square.CheckValidity()
		if err != nil {
			actual = false
		} else {
			actual = true
		}
		if actual != expected {
			t.Error(fmt.Sprintf("Expected %v, got %v", expected, actual))
		}
	}
}
