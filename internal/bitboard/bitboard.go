package bitboard

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"

	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

type Bitboard uint64

func NewBitboard(board uint64) Bitboard {
	return Bitboard(board)
}

func GetBit(board Bitboard, square sq.Square) bool {
	if !sq.OnBoard(square) {
		panic(fmt.Sprintf("tried to get bit off board: %d", square))
	}

	occ := (uint64(board) >> square) & 1

	// occ is a uint64, so we need to convert it to a bool
	return occ == 1
}

func SetBit(board Bitboard, square sq.Square) Bitboard {
	if !sq.OnBoard(square) {
		panic(fmt.Sprintf("tried to set bit off board: %d", square))
	}

	return Bitboard(uint64(board) | (1 << square))
}

func SetBits(board Bitboard, squares ...sq.Square) Bitboard {
	for _, square := range squares {
		board = SetBit(board, square)
	}

	return board
}

func ClearBit(board Bitboard, square sq.Square) Bitboard {
	if !sq.OnBoard(square) {
		panic(fmt.Sprintf("tried to clear bit off board: %d", square))
	}

	return board &^ (1 << square)
}

func LSBIndex(board Bitboard) sq.Square {
	if board == 0 {
		return sq.NoSquare
	}

	return sq.Square(CountBits(board&-board - 1))
}

func CountBits(board Bitboard) uint8 {
	return uint8(bits.OnesCount64(uint64(board)))
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
func SetOccupancy(index int, attackMask Bitboard) Bitboard {
	var ret Bitboard

	bitsInMask := CountBits(attackMask)

	for i := uint8(0); i < bitsInMask; i++ {
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
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}

	return binary.BigEndian.Uint64(b[:])
}

func GenerateSparseRandomUint64() uint64 {
	return RandomUint64() & RandomUint64() & RandomUint64() //nolint
}

// PrintBoard prints a bitboard to the console.
func PrintBoard(board Bitboard, output io.Writer) {
	if _, err := output.Write([]byte("\n")); err != nil {
		panic(err)
	}

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			// Convert rank and file into a square
			square := sq.Square(rank*8 + file)

			// Print the rank
			if file == 0 {
				if _, err := output.Write([]byte(fmt.Sprintf("%d  ", 8-rank))); err != nil {
					panic(err)
				}
			}

			// Check if the square is occupied
			occupied := GetBit(board, square)
			if occupied {
				if _, err := output.Write([]byte(fmt.Sprintf("%d ", 1))); err != nil {
					panic(err)
				}
			} else {
				if _, err := output.Write([]byte(fmt.Sprintf("%d ", 0))); err != nil {
					panic(err)
				}
			}
		}

		if _, err := output.Write([]byte("\n")); err != nil {
			panic(err)
		}
	}

	if _, err := output.Write([]byte("   a b c d e f g h")); err != nil {
		panic(err)
	}

	// Print the decimal representation
	if _, err := output.Write([]byte(fmt.Sprintf("\n   bitboard: %d\n", board))); err != nil {
		panic(err)
	}
}
