package tables

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/magic"
)

func TestMaskRookAttacks(t *testing.T) {
	var tests = map[int]uint64{
		sq.E4: 4521664529305600,    // central
		sq.G4: 18085034619584512,   // g-file
		sq.B5: 565159647117824,     // b-file
		sq.H4: 36170077829103616,   // h-file
		sq.A6: 282578808340736,     // a-file
		sq.E1: 7930856604974452736, // 1st rank
		sq.D7: 2260630401218048,    // 7th rank
		sq.H8: 36170086419038334,   // corner
		sq.A1: 9079539427579068672, // another corner
		sq.H2: 35607136465616896,   // another corner
	}

	for square, expected := range tests {
		actual := MaskRookAttacks(square)
		if uint64(actual) != expected {
			fmt.Println("Square", sq.Stringify(square))

			fmt.Println("Got")
			bb.PrintBoard(actual)

			fmt.Println("Expected")
			bb.PrintBoard(bb.Bitboard(expected))

			fmt.Println("")

			t.Errorf("Computing rook attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}

func TestRookAttacksOnTheFly(t *testing.T) {
	var blockers bb.Bitboard
	blockers = bb.SetBit(blockers, sq.D7)
	blockers = bb.SetBit(blockers, sq.D3)
	blockers = bb.SetBit(blockers, sq.F4)
	blockers = bb.SetBit(blockers, sq.B4)
	blockers = bb.SetBit(blockers, sq.A4)

	rookAttacks := RookAttacksOnTheFly(sq.D4, blockers)

	if rookAttacks != 9028156000256 {
		fmt.Println("Blockers:")
		bb.PrintBoard(blockers)

		fmt.Println("Rook attacks:")
		bb.PrintBoard(RookAttacksOnTheFly(sq.D4, blockers))
		t.Error("Expected 9028156000256, got ", rookAttacks)
	}
}

func TestLookupTableGivesCorrectMovesForRook(t *testing.T) {
	var data magic.Data
	if err := json.Unmarshal(magic.JsonData, &data); err != nil {
		panic(err)
	}

	table := populateRookAttackTables(data.Rook)

	testCases := []struct {
		square        int
		blockers      bb.Bitboard
		expectedMoves uint64
	}{
		{
			square:        sq.D4,
			blockers:      0,
			expectedMoves: 578722409201797128,
		},
		{
			square:        sq.D4,
			blockers:      bb.SetBit(0, sq.D7),
			expectedMoves: 578722409201797120,
		},
		{
			square:        sq.D4,
			blockers:      8796093024256,
			expectedMoves: 9857084688384,
		},
		{
			square:        sq.A1,
			blockers:      2342153281209368576,
			expectedMoves: 4467852305328242688,
		},
		{
			square:        sq.G8,
			blockers:      274877907074,
			expectedMoves: 275955859646,
		},
		{
			square:        sq.H1,
			blockers:      288239322568624128,
			expectedMoves: 8971311747122102272,
		},
	}

	for _, tt := range testCases {
		magicNum, err := strconv.ParseUint(data.Rook.Magics[tt.square].Magic, 16, 64)
		if err != nil {
			panic(err)
		}

		mask, err := strconv.ParseUint(data.Rook.Magics[tt.square].Mask, 16, 64)
		if err != nil {
			panic(err)
		}

		shift := data.Rook.Magics[tt.square].Shift
		index := (uint64(tt.blockers) & mask * magicNum) >> shift
		moves := table[tt.square][index]

		if uint64(moves) != tt.expectedMoves {
			fmt.Println("Blockers:")
			bb.PrintBoard(tt.blockers)
			fmt.Println("")

			fmt.Println("Expected moves:")
			bb.PrintBoard(bb.Bitboard(tt.expectedMoves))
			fmt.Println("")

			fmt.Println("Got moves:")
			bb.PrintBoard(moves)
			fmt.Println("")

			t.Error("For rook on square", sq.Stringify(tt.square), "expected", tt.expectedMoves, "got", moves)
		}
	}
}
