package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

var rookAttacks [64]uint64

func PopulateRookAttackTables() {
	for square := 0; square < 64; square++ {
		rookAttacks[square] = MaskRookAttacks(square)
	}
}

func MaskRookAttacks(square int) uint64 {
	var attackBoard uint64

	startRank := square / 8
	startFile := square % 8

	// North
	for rank := startRank - 1; rank > 0; rank-- {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+startFile)
	}

	// South
	for rank := startRank + 1; rank < 7; rank++ {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+startFile)
	}

	// East
	for file := startFile + 1; file < 7; file++ {
		attackBoard = bitboard.SetBit(attackBoard, startRank*8+file)
	}

	// West
	for file := startFile - 1; file > 0; file-- {
		attackBoard = bitboard.SetBit(attackBoard, startRank*8+file)
	}

	return attackBoard
}

// RookAttacksOnTheFly manually computes the possible squares a rook can attack
// depending on its position and a given blocker configuration.
func RookAttacksOnTheFly(square int, blockeres uint64) uint64 {
	var attackBoard uint64

	startRank := square / 8
	startFile := square % 8

	// North
	for rank := startRank - 1; rank >= 0; rank-- {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+startFile)
		if uint64(1)<<(rank*8+startFile)&blockeres != 0 {
			break
		}
	}

	// South
	for rank := startRank + 1; rank <= 7; rank++ {
		attackBoard = bitboard.SetBit(attackBoard, rank*8+startFile)
		if uint64(1)<<(rank*8+startFile)&blockeres != 0 {
			break
		}
	}

	// East
	for file := startFile + 1; file <= 7; file++ {
		attackBoard = bitboard.SetBit(attackBoard, startRank*8+file)
		if uint64(1)<<(startRank*8+file)&blockeres != 0 {
			break
		}
	}

	// West
	for file := startFile - 1; file >= 0; file-- {
		attackBoard = bitboard.SetBit(attackBoard, startRank*8+file)
		if uint64(1)<<(startRank*8+file)&blockeres != 0 {
			break
		}
	}

	return attackBoard
}
