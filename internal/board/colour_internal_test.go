package board

import (
	"testing"
)

func TestColourString(t *testing.T) {
	if Black.String() != "Black" {
		t.Errorf("Expected Black, got %s", Black.String())
	}

	if White.String() != "White" {
		t.Errorf("Expected White, got %s", White.String())
	}
}

func TestColourFromString(t *testing.T) {
	testCases := map[string]Colour{
		"Black": Black,
		"White": White,
		"black": Black,
		"white": White,
		"b":     Black,
		"w":     White,
		"foo":   Unknown,
	}

	for input, expected := range testCases {
		actual := ColourFromString(input)
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}
}
