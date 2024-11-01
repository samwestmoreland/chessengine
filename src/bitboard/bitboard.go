package bitboard

import "fmt"

func GetBit(board uint64, square int) bool {
	occ := (board >> square) & 1

	// occ is a uint64, so we need to convert it to a bool
	return occ == 1
}

func SetBit(board uint64, square int) uint64 {
	return board | (1 << square)
}

func ClearBit(board uint64, square int) uint64 {
	return board &^ (1 << square)
}

func CountBits(board uint64) int {
	count := 0
	for board != 0 {
		count++
		board &= board - 1
	}

	return count
}

func LSBIndex(board uint64) int {
	if board == 0 {
		return -1
	}

	return CountBits(board&-board - 1)
}

func SetOccupancy(index int, attackMask uint64) uint64 {
	var ret uint64

	bitsInMask := CountBits(attackMask)

	for i := 0; i < bitsInMask; i++ {
		sq := LSBIndex(attackMask)

		attackMask = ClearBit(attackMask, sq)

		if index&(1<<i) != 0 {
			ret |= (1 << sq)
		}
	}

	return ret
}

// print prints a bitboard to the console.
func PrintBoard(bitboard uint64) {
	fmt.Printf("\n")

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			// Convert rank and file into a square
			square := rank*8 + file

			// Print the rank
			if file == 0 {
				fmt.Printf("%d  ", 8-rank)
			}

			// Check if the square is occupied
			occupied := GetBit(bitboard, square)
			if occupied {
				fmt.Printf("%d ", 1)
			} else {
				fmt.Printf("%d ", 0)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println("   a b c d e f g h")

	// Print the decimal representation
	fmt.Printf("\n   bitboard: %d\n", bitboard)
}
