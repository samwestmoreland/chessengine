package main

import (
	"fmt"
	"os"

	"github.com/integrii/flaggy"
	"github.com/samwestmoreland/chessengine/src/position"
)

type colour string

const (
	white colour = "white"
	black colour = "black"
)

var version = "v0.0.0"

func main() {
	flaggy.SetName("Chess Engine")
	flaggy.SetDescription("A chess engine written in Go")
	flaggy.SetVersion(version)

	var fenFlag string
	flaggy.String(&fenFlag, "f", "fen", "A FEN string to parse")

	flaggy.Parse()
	fmt.Println("Starting chess engine")
	_, err := position.ParseFEN(fenFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return

}
