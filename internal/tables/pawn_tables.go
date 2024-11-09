// package tables
package tables

import (
	"github.com/samwestmoreland/chessengine/internal/bitboard"
	wb "github.com/samwestmoreland/chessengine/internal/colours"
)

// TODO: You can't ever have a white pawn on the 1st rank, so do we need to compute those?
func populatePawnAttackTables() [2][64]uint64 {
	var attacks [2][64]uint64

	for side := 0; side < 2; side++ {
		for square := 0; square < 64; square++ {
			attacks[side][square] = maskPawnAttacks(side, square)
		}
	}

	return attacks
}

func maskPawnAttacks(side, square int) uint64 {
	var attacks uint64

	board := bitboard.SetBit(0, square)

	if side == wb.White {
		attacks = maskWhitePawnAttacks(board, square)
	} else {
		attacks = maskBlackPawnAttacks(board, square)
	}

	return bitboard.ClearBit(attacks, square)
}

func maskBlackPawnAttacks(board uint64, square int) uint64 {
	if isAFile(square) {
		return board | (board << 9)
	}

	if isHFile(square) {
		return board | (board << 7)
	}

	return board | (board << 7) | (board << 9)
}

func maskWhitePawnAttacks(board uint64, square int) uint64 {
	if isAFile(square) {
		return board | (board >> 7)
	}

	if isHFile(square) {
		return board | (board >> 9)
	}

	return board | (board >> 7) | (board >> 9)
}
