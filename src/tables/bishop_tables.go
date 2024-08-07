package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

var bishopAttacks [64]uint64

func PopulateBishopAttackTables() {
	for square := 0; square < 64; square++ {
		bishopAttacks[square] = computeBishopAttacks(square)
	}
}

func computeBishopAttacks(square int) uint64 {
	var attackBoard uint64

	start_rank := square / 8
	start_file := square % 8

	// Bottom right
	for rank, file := start_rank+1, start_file+1; rank < 7 && file < 7; rank, file = rank+1, file+1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	// Top left
	for rank, file := start_rank-1, start_file-1; rank > 0 && file > 0; rank, file = rank-1, file-1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	// Top right
	for rank, file := start_rank-1, start_file+1; rank > 0 && file < 7; rank, file = rank-1, file+1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	// Bottom left
	for rank, file := start_rank+1, start_file-1; rank < 7 && file > 0; rank, file = rank+1, file-1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
	}

	return attackBoard
}
