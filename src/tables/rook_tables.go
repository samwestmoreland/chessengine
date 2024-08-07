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

	start_rank := square / 8
	start_file := square % 8

	// Vertical
	if start_file != 0 && start_file != 7 {
		for rank := 1; rank < 7; rank++ {
			if rank == start_rank {
				continue
			}

			attackBoard = bitboard.SetBit(attackBoard, rank*8+start_file)
		}
	}

	// Horizontal
	if start_rank != 0 && start_rank != 7 {
		for file := 1; file < 7; file++ {
			if file == start_file {
				continue
			}

			attackBoard = bitboard.SetBit(attackBoard, start_rank*8+file)
		}
	}

	return attackBoard
}
