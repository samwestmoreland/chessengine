// package tables
package tables

import "github.com/samwestmoreland/chessengine/src/bitboard"

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

// pawn_attacks [side][square]
// You can't ever have a white pawn on the 1st rank, so do we need to be computing those? Revisit later
var pawnAttacks [2][64]uint64

func populatePawnAttacks() {
	for side := 0; side < 2; side++ {
		for square := 0; square < 64; square++ {
			pawnAttacks[side][square] = computePawnAttacks(square, side)
		}
	}
}

func computePawnAttacks(square, side int) uint64 {
	var attacks uint64

	board := bitboard.SetBit(0, square)
	// bitboard.PrintBoard(board)

	// side == 0 is white
	if side == 0 {
		if isAFile(square) {
			attacks = board | (board >> 7)
		} else if isHFile(square) {
			attacks = board | (board >> 9)
		} else {
			attacks = board | (board >> 7) | (board >> 9)
		}
	} else {
		if isAFile(square) {
			attacks = board | (board << 9)
		} else if isHFile(square) {
			attacks = board | (board << 7)
		} else {
			attacks = board | (board << 7) | (board << 9)
		}
	}

	return bitboard.ClearBit(attacks, square)
}

func isAFile(square int) bool {
	return (1 << square & notAFile) == 0
}

func isHFile(square int) bool {
	return (1 << square & notHFile) == 0
}
