package tables

import (
	"strconv"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/magic"
)

func populateBishopAttackTables(data magic.BishopData) [64][]bb.Bitboard {
	var attacks [64][]bb.Bitboard

	for square := 0; square < 64; square++ {
		// Get magic data for this square
		magicNum, _ := strconv.ParseUint(data.Magics[square].Magic, 16, 64)
		shift := data.Magics[square].Shift

		// Create slice big enough for all possible indices
		tableSize := 1 << (64 - shift)
		attacks[square] = make([]bb.Bitboard, tableSize)

		// Populate this square's table with all possible attack patterns
		mask := MaskBishopAttacks(square)
		numBlockers := bb.CountBits(mask) // how many relevant squares

		// For each possible blocker configuration...
		for i := 0; i < (1 << numBlockers); i++ {
			blockers := bb.SetOccupancy(i, mask)
			// Calculate actual moves for this blocker pattern
			moves := BishopAttacksOnTheFly(square, blockers)
			// Calculate index using magic
			index := (uint64(blockers) * magicNum) >> shift
			// Store moves at this index
			attacks[square][index] = moves
		}
	}

	return attacks
}

func MaskBishopAttacks(square int) bb.Bitboard {
	var attackBoard bb.Bitboard

	startRank := square / 8
	startFile := square % 8

	// Bottom right
	for rank, file := startRank+1, startFile+1; rank < 7 && file < 7; rank, file = rank+1, file+1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
	}

	// Top left
	for rank, file := startRank-1, startFile-1; rank > 0 && file > 0; rank, file = rank-1, file-1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
	}

	// Top right
	for rank, file := startRank-1, startFile+1; rank > 0 && file < 7; rank, file = rank-1, file+1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
	}

	// Bottom left
	for rank, file := startRank+1, startFile-1; rank < 7 && file > 0; rank, file = rank+1, file-1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
	}

	return attackBoard
}

func BishopAttacksOnTheFly(square int, blockers bb.Bitboard) bb.Bitboard {
	var attackBoard bb.Bitboard

	startRank := square / 8
	startFile := square % 8

	// Bottom right
	for rank, file := startRank+1, startFile+1; rank <= 7 && file <= 7; rank, file = rank+1, file+1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&uint64(blockers) != 0 {
			break
		}
	}

	// Top left
	for rank, file := startRank-1, startFile-1; rank >= 0 && file >= 0; rank, file = rank-1, file-1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&uint64(blockers) != 0 {
			break
		}
	}

	// Top right
	for rank, file := startRank-1, startFile+1; rank >= 0 && file <= 7; rank, file = rank-1, file+1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&uint64(blockers) != 0 {
			break
		}
	}

	// Bottom left
	for rank, file := startRank+1, startFile-1; rank <= 7 && file >= 0; rank, file = rank+1, file-1 {
		attackBoard = bb.SetBit(attackBoard, rank*8+file)
		if uint64(1)<<(rank*8+file)&uint64(blockers) != 0 {
			break
		}
	}

	return attackBoard
}
