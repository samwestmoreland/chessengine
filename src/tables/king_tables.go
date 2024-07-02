// package tables
package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

var kingAttacks [64]uint64

func PopulateKingAttackTables() {
	for square := 0; square < 64; square++ {
		kingAttacks[square] = computeKingAttacks(square)
	}
}

func computeKingAttacks(square int) uint64 {
	pieceBoard := bitboard.SetBit(0, square)

	var attackBoard uint64

	if pieceBoard>>8 != 0 {
		attackBoard |= (pieceBoard >> 8)
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
