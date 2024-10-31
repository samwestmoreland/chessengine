package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

var bishopAttacks [64]uint64

func PopulateBishopAttackTables() {
	for square := 0; square < 64; square++ {
		bishopAttacks[square] = maskBishopAttacks(square)
	}
}

func maskBishopAttacks(square int) uint64 {
	var attackBoard uint64

	startRank := square / 8
	startFile := square % 8

	// Bottom right
	for rank, file := startRank+1, startFile+1; rank < 7 && file < 7; rank, file = rank+1, file+1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	// Top left
	for rank, file := startRank-1, startFile-1; rank > 0 && file > 0; rank, file = rank-1, file-1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	// Top right
	for rank, file := startRank-1, startFile+1; rank > 0 && file < 7; rank, file = rank-1, file+1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	// Bottom left
	for rank, file := startRank+1, startFile-1; rank < 7 && file > 0; rank, file = rank+1, file-1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	return attackBoard
}
