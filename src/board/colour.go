package board

import (
	"strings"
)

type Colour int

const (
	White Colour = iota
	Black
	Unknown
)

func (c Colour) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	case Unknown:
		return "Unknown"
	}

	return "Unknown"
}

func ColourFromString(colStr string) Colour {
	colStr = strings.ToLower(colStr)
	switch colStr {
	case "white":
		return White
	case "w":
		return White
	case "black":
		return Black
	case "b":
		return Black
	}

	return Unknown
}
