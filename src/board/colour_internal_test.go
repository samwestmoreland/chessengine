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
	if ColourFromString("Black") != Black {
		t.Errorf("Expected Black, got %s", ColourFromString("Black"))
	}

	if ColourFromString("White") != White {
		t.Errorf("Expected White, got %s", ColourFromString("White"))
	}

	if ColourFromString("Invalid") != Unknown {
		t.Errorf("Expected InvalidColour, got %s", ColourFromString("Invalid"))
	}
}
