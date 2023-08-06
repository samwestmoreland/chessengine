package position

import (
	"fmt"
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

func TestParseInvalidFENColour(t *testing.T) {
	t.Log("Testing invalid FEN")

	input := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR p KQkq - 0 1"

	_, err := ParseFEN(input)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestParseInvalidFENCastling(t *testing.T) {
	t.Log("Testing invalid FEN")

	input := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	_, err := ParseFEN(input)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseInvalidFENEnPassant(t *testing.T) {
	t.Log("Testing invalid FEN")

	validEnPassantSquares := []string{
		"a3", "a6", "b3", "b6",
		"c3", "c6", "d3", "d6",
		"e3", "e6", "f3", "f6",
		"g3", "g6", "h3", "h6"}

	for _, square := range validEnPassantSquares {
		fen := fmt.Sprintf("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq %s 0 1", square)

		_, err := ParseFEN(fen)
		if err != nil {
			t.Error(err)
		}
	}

	invalidEnPassantSquares := []string{"&2", "a9", "i3", "a", "3", "-1", "z", "zzz", "-2", "a0"}
	for _, square := range invalidEnPassantSquares {
		fen := fmt.Sprintf("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq %s 0 1", square)

		_, err := ParseFEN(fen)
		if err == nil {
			t.Error("Expected error")
		}
	}
}
