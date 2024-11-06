package tables

import (
	"math/bits"
	"strconv"

	"github.com/samwestmoreland/chessengine/magic"
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

func populateBishopAttackTables(data magic.BishopData) [64][]uint64 {
	var attacks [64][]uint64

	for square := 0; square < 64; square++ {
		// Get magic data for this square
		magicNum, _ := strconv.ParseUint(data.Magics[square].Magic, 16, 64)
		shift := data.Magics[square].Shift

		// Create slice big enough for all possible indices
		tableSize := 1 << (64 - shift)
		attacks[square] = make([]uint64, tableSize)

		// Populate this square's table with all possible attack patterns
		mask := MaskBishopAttacks(square)
		numBlockers := bits.OnesCount64(mask) // how many relevant squares

		// For each possible blocker configuration...
		for i := 0; i < (1 << numBlockers); i++ {
			blockers := bitboard.SetOccupancy(i, mask)
			// Calculate actual moves for this blocker pattern
			moves := BishopAttacksOnTheFly(square, blockers)
			// Calculate index using magic
			index := (blockers * magicNum) >> shift
			// Store moves at this index
			attacks[square][index] = moves
		}
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
