package tables

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/samwestmoreland/chessengine/magic"
	"github.com/samwestmoreland/chessengine/src/bitboard"
	sq "github.com/samwestmoreland/chessengine/src/squares"
)

func TestMaskRookAttacks(t *testing.T) {
	var rookTestCases = map[int]uint64{
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

	for square, expected := range rookTestCases {
		actual := MaskRookAttacks(square)
		if actual != expected {
			fmt.Println("Square", sq.Stringify(square))

			fmt.Println("Got")
			bitboard.PrintBoard(actual)

			fmt.Println("Expected")
			bitboard.PrintBoard(expected)

			fmt.Println("")

			t.Errorf("Computing rook attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}

func TestRookAttacksOnTheFly(t *testing.T) {
	var blockers uint64
	blockers = bitboard.SetBit(blockers, sq.D7)
	blockers = bitboard.SetBit(blockers, sq.D3)
	blockers = bitboard.SetBit(blockers, sq.F4)
	blockers = bitboard.SetBit(blockers, sq.B4)
	blockers = bitboard.SetBit(blockers, sq.A4)

	rookAttacks := RookAttacksOnTheFly(sq.D4, blockers)

	if rookAttacks != 9028156000256 {
		fmt.Println("Blockers:")
		bitboard.PrintBoard(blockers)

		fmt.Println("Rook attacks:")
		bitboard.PrintBoard(RookAttacksOnTheFly(sq.D4, blockers))
		t.Error("Expected 9028156000256, got ", rookAttacks)
	}
}

func TestLookupTableGivesCorrectMoves(t *testing.T) {
	var data magic.Data
	if err := json.Unmarshal(magic.JsonData, &data); err != nil {
		panic(err)
	}

	table := populateRookAttackTables(data.Rook)

	testCases := []struct {
		square        int
		blockers      uint64
		expectedMoves uint64
	}{
		{
			square:        sq.D4,
			blockers:      0,
			expectedMoves: 578722409201797128,
		},
		{
			square:        sq.D4,
			blockers:      bitboard.SetBit(0, sq.D7),
			expectedMoves: 578722409201797120,
		},
		{
			square:        sq.D4,
			blockers:      bitboard.SetBit(0, sq.D7) | bitboard.SetBit(0, sq.D3),
			expectedMoves: 9857084688384,
		},
		{
			square:        sq.A1,
			blockers:      bitboard.SetBit(0, sq.A2) | bitboard.SetBit(0, sq.F1) | bitboard.SetBit(0, sq.H2),
			expectedMoves: 4467852305328242688,
		},
		{
			square:        sq.G8,
			blockers:      bitboard.SetBit(0, sq.G4) | bitboard.SetBit(0, sq.B8) | bitboard.SetBit(0, sq.H8),
			expectedMoves: 275955859646,
		},
	}

	for _, tt := range testCases {
		magicNum, err := strconv.ParseUint(data.Rook.Magics[tt.square].Magic, 16, 64)
		if err != nil {
			panic(err)
		}

		shift := data.Rook.Magics[tt.square].Shift
		index := (tt.blockers * magicNum) >> shift
		moves := table[tt.square][index]

		if moves != tt.expectedMoves {
			fmt.Println("Blockers:")
			bitboard.PrintBoard(tt.blockers)
			fmt.Println("")

			fmt.Println("Got moves:")
			bitboard.PrintBoard(moves)
			fmt.Println("")

			fmt.Println("Expected moves:")
			bitboard.PrintBoard(tt.expectedMoves)
			fmt.Println("")

			t.Error("Expected", tt.expectedMoves, "got", moves)
		}
	}
}
