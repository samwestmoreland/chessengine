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
		t.Error(err)
	}
}

func TestParseInvalidFENEnPassant(t *testing.T) {
	t.Log("Testing invalid FEN")
	input := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Qkq j3 0 1"
	_, err := ParseFEN(input)
	if err == nil {
		t.Error("Expected error")
	}
}
