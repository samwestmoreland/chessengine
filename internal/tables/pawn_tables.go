// package tables
package tables

import (
	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

// TODO: You can't ever have a white pawn on the 1st rank, so do we need to compute those?
func populatePawnAttackTables() [2][64]bb.Bitboard {
	var attacks [2][64]bb.Bitboard

	for side := uint8(0); side < 2; side++ {
		for square := uint8(0); square < 64; square++ {
			attacks[side][square] = maskPawnAttacks(side, sq.Square(square))
		}
	}

	return attacks
}

func maskPawnAttacks(side uint8, square sq.Square) bb.Bitboard {
	var attacks bb.Bitboard

	board := bb.SetBit(0, square)

	if side == 0 {
		attacks = maskWhitePawnAttacks(board, square)
	} else {
		attacks = maskBlackPawnAttacks(board, square)
	}

	return bb.ClearBit(attacks, square)
}

func maskBlackPawnAttacks(board bb.Bitboard, square sq.Square) bb.Bitboard {
	if isAFile(square) {
		return board | (board << 9)
	}

	if isHFile(square) {
		return board | (board << 7)
	}

	return board | (board << 7) | (board << 9)
}

func maskWhitePawnAttacks(board bb.Bitboard, square sq.Square) bb.Bitboard {
	if isAFile(square) {
		return board | (board >> 7)
	}

	if isHFile(square) {
		return board | (board >> 9)
	}

	return board | (board >> 7) | (board >> 9)
}
