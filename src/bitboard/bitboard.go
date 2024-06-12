package bitboard

import "fmt"

// print prints a bitboard to the console.
func print(bitboard uint64) {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			// Convert rank and file into a square
			square := rank*8 + file

			// Check if the square is occupied
			occupied := (bitboard >> square) & 1
			fmt.Printf(" %d ", occupied)
		}
		fmt.Printf("\n")
	}
}
