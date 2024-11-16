// package tables
package tables

import (
	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

func populateKingAttackTables() [64]bb.Bitboard {
	var attacks [64]bb.Bitboard

	for square := uint8(0); square < 64; square++ {
		attacks[square] = maskKingAttacks(sq.Square(square))
	}

	return attacks
}

func maskKingAttacks(square sq.Square) bb.Bitboard {
	pieceBoard := bb.SetBit(0, square)

	var attackBoard bb.Bitboard

	if pieceBoard>>8 != 0 {
		attackBoard = bb.Bitboard(uint64(attackBoard) | (uint64(pieceBoard) >> 8))
	}

	if pieceBoard>>9 != 0 && pieceBoard>>9&notHFile != 0 {
		attackBoard |= (pieceBoard >> 9)
	}

	if (pieceBoard>>7) != 0 && (pieceBoard>>7)&notAFile != 0 {
		attackBoard |= (pieceBoard >> 7)
	}

	if (pieceBoard>>1) != 0 && pieceBoard>>1&notHFile != 0 {
		attackBoard |= (pieceBoard >> 1)
	}

	if pieceBoard<<8 != 0 {
		attackBoard |= (pieceBoard << 8)
	}

	if pieceBoard<<9 != 0 && pieceBoard<<9&notAFile != 0 {
		attackBoard |= (pieceBoard << 9)
	}

	if (pieceBoard<<7) != 0 && (pieceBoard<<7)&notHFile != 0 {
		attackBoard |= (pieceBoard << 7)
	}

	if (pieceBoard<<1) != 0 && pieceBoard<<1&notAFile != 0 {
		attackBoard |= (pieceBoard << 1)
	}

	return attackBoard
}
