// package tables
package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

var knightAttacks [64]uint64

func PopulateKnightAttackTables() {
	for square := 0; square < 64; square++ {
		knightAttacks[square] = computeKnightAttacks(square)
	}
}

func computeKnightAttacks(square int) uint64 {
	board := bitboard.SetBit(0, square)

	if isAFile(square) {
		board |= (board << 10) | (board << 17) |
			(board >> 6) | (board >> 15)

		return bitboard.ClearBit(board, square)
	}

	if isHFile(square) {
		board |= (board >> 10) | (board >> 17) |
			(board << 6) | (board << 15)

		return bitboard.ClearBit(board, square)
	}

	if isABFile(square) {
		board |= (board >> 6) | (board >> 15) | (board >> 17) |
			(board << 10) | (board << 15) | (board << 17)

		return bitboard.ClearBit(board, square)
	}

	if isGHFile(square) {
		board |= (board << 6) | (board << 15) | (board << 17) |
			(board >> 10) | (board >> 15) | (board >> 17)

		return bitboard.ClearBit(board, square)
	}

	board |= (board << 6) | (board << 10) | (board << 15) | (board << 17) |
		(board >> 10) | (board >> 6) | (board >> 15) | (board >> 17)

	return bitboard.ClearBit(board, square)
}
