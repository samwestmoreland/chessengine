package bitboard

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/bits"
)

func GetBit(board uint64, square int) bool {
	occ := (board >> square) & 1

	// occ is a uint64, so we need to convert it to a bool
	return occ == 1
}

func SetBit(board uint64, square int) uint64 {
	return board | (1 << square)
}

func SetBits(board uint64, squares ...int) uint64 {
	for _, square := range squares {
		board = SetBit(board, square)
	}

	return board
}

func ClearBit(board uint64, square int) uint64 {
	return board &^ (1 << square)
}

func LSBIndex(board uint64) int {
	if board == 0 {
		return -1
	}

	return bits.OnesCount64(board&-board - 1)
}

// SetOccupancy sets each bit on the attack mask to 1 or 0.
//
// E.g. consider index 9, and an attack mask for a rook on d4:
// 9    = 0 0000 1001
//
// There are 10 set bits in the mask, so we represent the index with 10 bits.
// We go to the least significant bit, which in this case is d7:
//
// index&(1<<0) = 0 0000 0001
// 9            = 0 0000 1001
// -> set bit on d7
//
// then in the next iteration, we go to d6:
// index&(1<<1) = 0 0000 0010
// 9            = 0 0000 1001
// -> don't set bit on d6
//
// and so on...
func SetOccupancy(index int, attackMask uint64) uint64 {
	var ret uint64

	bitsInMask := bits.OnesCount64(attackMask)

	for i := 0; i < bitsInMask; i++ {
		sq := LSBIndex(attackMask)

		attackMask = ClearBit(attackMask, sq)

		// Set the bit on the occupancy board
		if index&(1<<i) != 0 {
			ret |= (1 << sq)
		}
	}

	return ret
}

func RandomUint64() uint64 {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err) // or handle error appropriately
	}

	return binary.BigEndian.Uint64(b[:])
}

func GenerateSparseRandomUint64() uint64 {
	return RandomUint64() & RandomUint64() & RandomUint64()
}

// PrintBoard prints a bitboard to the console.
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
