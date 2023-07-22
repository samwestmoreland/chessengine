package main

import (
	"fmt"
	"os"
	"strings"
)

type position string

// we want a colour type that is an enum of white and black
type colour string

const (
	white colour = "white"
	black colour = "black"
)

func main() {
	fmt.Println("Starting chess engine")
	fmt.Println("")
	fmt.Println("Please enter a position in FEN notation")

	// read in the position
	// check that it is valid
	var pos position
	fmt.Scanln(&pos)
	if !pos.isValid() {
		fmt.Println("Invalid position")
		os.Exit(1)
	}

	// print out the position
	fmt.Println(pos)
	return
}

func (p position) isValid() bool {
	// check that there are 8 rows
	rows := p.splitRows()
	if len(rows) != 8 {
		return false
	}

	return true
}

func (p position) splitRows() []string {
	// p is a string with some slashes in it
	// we want to split it into 8 strings
	// convert p, a position, to a string
	str := string(p)

	rows := make([]string, 8)

	// split the string into 8 strings
	// split on the slash
	rows = strings.Split(str, "/")

	return rows
}

func (p position) nextColourToMove() string {
	return ""
}
