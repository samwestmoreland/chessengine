package tables

import (
	"log"
	"strconv"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/magic"
)

func populateBishopAttackTables(data magic.BishopData) [64][]bb.Bitboard {
	var attacks [64][]bb.Bitboard

	for square := range uint8(64) {
		// Get magic data for this square
		magicNum, err := strconv.ParseUint(data.Magics[square].Magic, 16, 64)
		if err != nil {
			log.Fatal(err)
		}

		shift := data.Magics[square].Shift

		// Create slice big enough for all possible indices
		tableSize := 1 << (64 - shift)
		attacks[square] = make([]bb.Bitboard, tableSize)

		// Populate this square's table with all possible attack patterns
		mask := MaskBishopAttacks(sq.Square(square))
		numBlockers := bb.CountBits(mask) // how many relevant squares

		// For each possible blocker configuration...
		for i := range 1 << numBlockers {
			blockers := bb.SetOccupancy(i, mask)
			// Calculate actual moves for this blocker pattern
			moves := BishopAttacksOnTheFly(sq.Square(square), blockers)
			// Calculate index using magic
			index := (uint64(blockers) * magicNum) >> shift
			// Store moves at this index
			attacks[square][index] = moves
		}
	}

	return attacks
}

func MaskBishopAttacks(square sq.Square) bb.Bitboard {
	var attackBoard bb.Bitboard

	startRank := int(square / 8)
	startFile := int(square % 8)

	directions := [4][2]int{
		{1, 1},   // bottom right
		{1, -1},  // bottom left
		{-1, 1},  // top right
		{-1, -1}, // top left
	}

	for _, dir := range directions {
		rankDelta := dir[0]
		fileDelta := dir[1]

		for rank, file := startRank+rankDelta, startFile+fileDelta; rank > 0 && rank < 7 &&
			file > 0 && file < 7; rank, file = rank+rankDelta, file+fileDelta {
			square := sq.Square(byte(rank*8 + file))
			attackBoard = bb.SetBit(attackBoard, square)
		}
	}

	return attackBoard
}

func BishopAttacksOnTheFly(square sq.Square, blockers bb.Bitboard) bb.Bitboard {
	var attackBoard bb.Bitboard

	startRank := int(square / 8)
	startFile := int(square % 8)

	directions := [4][2]int{
		{1, 1},   // bottom right
		{1, -1},  // bottom left
		{-1, 1},  // top right
		{-1, -1}, // top left
	}

	for _, dir := range directions {
		rankDelta := dir[0]
		fileDelta := dir[1]

		for rank, file := startRank+rankDelta, startFile+fileDelta; rank >= 0 && rank <= 7 &&
			file >= 0 && file <= 7; rank, file = rank+rankDelta, file+fileDelta {
			square := sq.Square(byte(rank*8 + file))
			attackBoard = bb.SetBit(attackBoard, square)

			if uint64(1)<<(rank*8+file)&uint64(blockers) != 0 {
				break
			}
		}
	}

	return attackBoard
}
