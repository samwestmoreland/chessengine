// package tables
package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
	wb "github.com/samwestmoreland/chessengine/src/colours"
)

// notAFile is const represeting the board:
//
//    8  0 1 1 1 1 1 1 1
//    7  0 1 1 1 1 1 1 1
//    6  0 1 1 1 1 1 1 1
//    5  0 1 1 1 1 1 1 1
//    4  0 1 1 1 1 1 1 1
//    3  0 1 1 1 1 1 1 1
//    2  0 1 1 1 1 1 1 1
//    1  0 1 1 1 1 1 1 1
//       a b c d e f g h
//

const notAFile uint64 = 18374403900871474942

// notHFile is const represeting the board:
//
//    8  1 1 1 1 1 1 1 0
//    7  1 1 1 1 1 1 1 0
//    6  1 1 1 1 1 1 1 0
//    5  1 1 1 1 1 1 1 0
//    4  1 1 1 1 1 1 1 0
//    3  1 1 1 1 1 1 1 0
//    2  1 1 1 1 1 1 1 0
//    1  1 1 1 1 1 1 1 0
//       a b c d e f g h
//

const notHFile uint64 = 9187201950435737471

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

func isAFile(square int) bool {
	return (1 << square & notAFile) == 0
}

func isHFile(square int) bool {
	return (1 << square & notHFile) == 0
}
