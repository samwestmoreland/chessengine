// package tables
package tables

import (
	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

func populateKnightAttackTables() [64]bb.Bitboard {
	var attacks [64]bb.Bitboard

	for square := range 64 {
		attacks[square] = maskKnightAttacks(sq.Square(byte(square)))
	}

	return attacks
}

func maskKnightAttacks(square sq.Square) bb.Bitboard {
	board := bb.SetBit(0, square)

	if isAFile(square) {
		board |= (board << 10) | (board << 17) |
			(board >> 6) | (board >> 15)

		return bb.ClearBit(board, square)
	}

	if isHFile(square) {
		board |= (board >> 10) | (board >> 17) |
			(board << 6) | (board << 15)

		return bb.ClearBit(board, square)
	}

	if isABFile(square) {
		board |= (board >> 6) | (board >> 15) | (board >> 17) |
			(board << 10) | (board << 15) | (board << 17)

		return bb.ClearBit(board, square)
	}

	if isGHFile(square) {
		board |= (board << 6) | (board << 15) | (board << 17) |
			(board >> 10) | (board >> 15) | (board >> 17)

		return bb.ClearBit(board, square)
	}

	board |= (board << 6) | (board << 10) | (board << 15) | (board << 17) |
		(board >> 10) | (board >> 6) | (board >> 15) | (board >> 17)

	return bb.ClearBit(board, square)
}
