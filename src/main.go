package main

import (
	"fmt"
	"os"

	"github.com/samwestmoreland/chessengine/src/position"
)

type colour string

const (
	white colour = "white"
	black colour = "black"
)

func main() {
	fmt.Println("Starting chess engine")
	fmt.Println("")
	fmt.Println("Please enter a position in FEN notation")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = position.ParseFEN(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return

}
