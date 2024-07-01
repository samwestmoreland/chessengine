// package tables
package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
	wb "github.com/samwestmoreland/chessengine/src/colours"
)

// pawnAttacks [side][square]
// You can't ever have a white pawn on the 1st rank, so do we need to be
// computing those? Revisit later.
var pawnAttacks [2][64]uint64

func PopulatePawnAttackTables() {
	for side := 0; side < 2; side++ {
		for square := 0; square < 64; square++ {
			pawnAttacks[side][square] = computePawnAttacks(side, square)
		}
	}
}

func computePawnAttacks(side, square int) uint64 {
	var attacks uint64

	board := bitboard.SetBit(0, square)

	if side == wb.White {
		attacks = computeWhitePawnAttacks(board, square)
	} else {
		attacks = computeBlackPawnAttacks(board, square)
	}

	return bitboard.ClearBit(attacks, square)
}

func computeBlackPawnAttacks(board uint64, square int) uint64 {
	if isAFile(square) {
		return board | (board << 9)
	}

	if isHFile(square) {
		return board | (board << 7)
	}

	return board | (board << 7) | (board << 9)
}

func computeWhitePawnAttacks(board uint64, square int) uint64 {
	if isAFile(square) {
		return board | (board >> 7)
	}

	if isHFile(square) {
		return board | (board >> 9)
	}

	return board | (board >> 7) | (board >> 9)
}
