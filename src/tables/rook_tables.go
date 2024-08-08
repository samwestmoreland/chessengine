package tables

import (
	"github.com/samwestmoreland/chessengine/src/bitboard"
)

var rookAttacks [64]uint64

func PopulateRookAttackTables() {
	for square := 0; square < 64; square++ {
		rookAttacks[square] = computeRookAttacks(square)
	}
}

func computeRookAttacks(square int) uint64 {
	var attackBoard uint64

	startRank := square / 8
	startFile := square % 8

	// Vertical
	if startFile != 0 && startFile != 7 {
		for rank := 1; rank < 7; rank++ {
			if rank == startRank {
				continue
			}

			attackBoard = bitboard.SetBit(attackBoard, rank*8+startFile)
		}
	}

	// Horizontal
	if startRank != 0 && startRank != 7 {
		for file := 1; file < 7; file++ {
			if file == startFile {
				continue
			}

			attackBoard = bitboard.SetBit(attackBoard, startRank*8+file)
		}
	}

	return attackBoard
}
