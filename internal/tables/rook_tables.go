package tables

import (
	"strconv"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/magic"
)

// populateRookAttackTables generates a lookup table for rook attacks. Each square on the board has
// its own hash table, which can be indexed into by hashing the blocker configuration like so:
//
// index = (blockerConfig * magicNumber) >> someShift
//
// where magicNumber is a magic number that has been pre-calculated and stored in the magic_data
// directory, and someShift is a bit shift that has also been pre-calculated and is stored alongside
// the magic number.
func populateRookAttackTables(data magic.RookData) [64][]bb.Bitboard {
	var attacks [64][]bb.Bitboard

	for square := uint8(0); square < 64; square++ {
		// Get magic data for this square
		magicNum, _ := strconv.ParseUint(data.Magics[square].Magic, 16, 64)
		shift := data.Magics[square].Shift

		// Create slice big enough for all possible indices
		tableSize := 1 << (64 - shift)
		attacks[square] = make([]bb.Bitboard, tableSize)

		// Populate this square's table with all possible attack patterns
		mask := MaskRookAttacks(sq.Square(square))
		numBlockers := bb.CountBits(mask) // how many relevant squares

		// For each possible blocker configuration...
		for i := 0; i < (1 << numBlockers); i++ {
			blockers := bb.SetOccupancy(i, mask)
			// Calculate actual moves for this blocker pattern
			moves := RookAttacksOnTheFly(sq.Square(square), blockers)
			// Calculate index using magic
			index := (uint64(blockers) * magicNum) >> shift
			// Store moves at this index
			attacks[square][index] = moves
		}
	}

	return attacks
}

// MaskRookAttacks generates a bitmask for all possible squares that a rook can attack from a given
// square.
func MaskRookAttacks(square sq.Square) bb.Bitboard {
	var attackBoard bb.Bitboard

	startRank := square / 8
	startFile := square % 8

	// North
	for rank := startRank - 1; rank > 0 && rank < 7; rank-- {
		attackBoard = bb.SetBit(attackBoard, rank*8+startFile)
	}

	// South
	for rank := startRank + 1; rank > 0 && rank < 7; rank++ {
		attackBoard = bb.SetBit(attackBoard, rank*8+startFile)
	}

	// East
	for file := startFile + 1; file > 0 && file < 7; file++ {
		attackBoard = bb.SetBit(attackBoard, startRank*8+file)
	}

	// West
	for file := startFile - 1; file > 0 && file < 7; file-- {
		attackBoard = bb.SetBit(attackBoard, startRank*8+file)
	}

	return attackBoard
}

// RookAttacksOnTheFly manually computes the possible squares a rook can attack
// depending on its position and a given blocker configuration.
func RookAttacksOnTheFly(square sq.Square, blockers bb.Bitboard) bb.Bitboard {
	var attackBoard bb.Bitboard

	startRank := square / 8
	startFile := square % 8

	// North
	for rank := startRank - 1; rank <= 7; rank-- {
		attackBoard = bb.SetBit(attackBoard, rank*8+startFile)

		if bb.Bitboard(1)<<(rank*8+startFile)&blockers != 0 {
			break
		}
	}

	// South
	for rank := startRank + 1; rank <= 7; rank++ {
		attackBoard = bb.SetBit(attackBoard, rank*8+startFile)

		if bb.Bitboard(1)<<(rank*8+startFile)&blockers != 0 {
			break
		}
	}

	// East
	for file := startFile + 1; file <= 7; file++ {
		attackBoard = bb.SetBit(attackBoard, startRank*8+file)

		if bb.Bitboard(1)<<(startRank*8+file)&blockers != 0 {
			break
		}
	}

	// West
	for file := startFile - 1; file <= 7; file-- {
		attackBoard = bb.SetBit(attackBoard, startRank*8+file)

		if bb.Bitboard(1)<<(startRank*8+file)&blockers != 0 {
			break
		}
	}

	return attackBoard
}
