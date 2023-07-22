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

	var pos position.Position
	fmt.Scanln(&pos)
	if !pos.IsValid() {
		fmt.Println("Invalid position")
		os.Exit(1)
	}

	// print out the position
	fmt.Println(pos)
	return
}
