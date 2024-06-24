// package tables
package tables

import "github.com/samwestmoreland/chessengine/src/bitboard"

// const not_a_file unit64 =

// const not_h_file unit64 =

// pawn_attacks [side][square]
var pawnAttacks [2][64]uint64

func MaskPawnAttacks(square, side int) uint64 {
	board := bitboard.SetBit(0, square)
	bitboard.PrintBoard(board)
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			// square := rank*8 + file
		}
	}

	return 0
}
