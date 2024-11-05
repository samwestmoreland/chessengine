package tables

import (
	"github.com/samwestmoreland/chessengine/magic"
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

func populateBishopAttackTables(magic.BishopData) [64][512]uint64 {
	var attacks [64][512]uint64

	for square := 0; square < 64; square++ {
	}

	return attacks
}

func MaskBishopAttacks(square int) uint64 {
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

func BishopAttacksOnTheFly(square int, blockers uint64) uint64 {
	var attackBoard uint64

	startRank := square / 8
	startFile := square % 8

	// Bottom right
	for rank, file := startRank+1, startFile+1; rank <= 7 && file <= 7; rank, file = rank+1, file+1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&blockers != 0 {
			break
		}
	}

	// Top left
	for rank, file := startRank-1, startFile-1; rank >= 0 && file >= 0; rank, file = rank-1, file-1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&blockers != 0 {
			break
		}
	}

	// Top right
	for rank, file := startRank-1, startFile+1; rank >= 0 && file <= 7; rank, file = rank-1, file+1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&blockers != 0 {
			break
		}
	}

	// Bottom left
	for rank, file := startRank+1, startFile-1; rank <= 7 && file >= 0; rank, file = rank+1, file-1 {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&blockers != 0 {
			break
		}
	}

	return attackBoard
}
