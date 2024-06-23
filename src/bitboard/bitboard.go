package bitboard

import "fmt"

func getBit(board uint64, square int) bool {
	occ := (board >> square) & 1

	// occ is a uint64, so we need to convert it to a bool
	return occ == 1
}

// print prints a bitboard to the console.
func printBoard(bitboard uint64) {
	fmt.Printf("\n")

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			// Convert rank and file into a square
			square := rank*8 + file

			// Print the rank
			if file == 0 {
				fmt.Printf("%d ", 8-rank)
			}

			// Check if the square is occupied
			occupied := getBit(bitboard, square)
			if occupied {
				fmt.Printf("%d ", 1)
			} else {
				fmt.Printf("%d ", 0)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println("  A B C D E F G H")

}
