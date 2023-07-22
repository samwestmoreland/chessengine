package position

import (
	"testing"
)

func TestParseValidFEN(t *testing.T) {
	t.Log("Testing valid FEN")
	input := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	_, err := ParseFEN(input)
	if err != nil {
		t.Error(err)
	}
}
