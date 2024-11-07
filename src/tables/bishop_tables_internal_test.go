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

var bishopTestCases = map[int]uint64{
	sq.E4: 19184279556981248, // central
	sq.G4: 4538784537380864,  // g-file
	sq.B5: 4512412900526080,  // b-file
	sq.H4: 9077569074761728,  // h-file
	sq.A6: 4512412933816832,  // a-file
	sq.E1: 11333774449049600, // 1st rank
	sq.D7: 275449643008,      // 7th rank
	sq.H8: 567382630219776,   // corner
	sq.A1: 567382630219776,   // another corner
	sq.H2: 70506452091904,    // another corner
}

func TestMaskBishopAttacks(t *testing.T) {
	for square, expected := range bishopTestCases {
		actual := MaskBishopAttacks(square)
		if actual != expected {
			bitboard.PrintBoard(actual)
			t.Errorf("Computing bishop attacks for %s, expected %d, got %d", sq.Stringify(square), expected, actual)
		}
	}
}

func TestBishopAttacksOnTheFly(t *testing.T) {
	var blockers uint64
	blockers = bitboard.SetBit(blockers, sq.B6)
	blockers = bitboard.SetBit(blockers, sq.G7)
	blockers = bitboard.SetBit(blockers, sq.E3)
	blockers = bitboard.SetBit(blockers, sq.B2)

	bishopAttacks := BishopAttacksOnTheFly(sq.D4, blockers)

	if bishopAttacks != 584940523765760 {
		fmt.Println("Blockers:")
		bitboard.PrintBoard(blockers)

		fmt.Println("Bishop attacks on the fly:")
		bitboard.PrintBoard(bishopAttacks)
		t.Error("Expected 584940523765760, got ", bishopAttacks)
	}
}

func TestLookupTableGivesCorrectMovesForBishop(t *testing.T) {
	var data magic.Data
	if err := json.Unmarshal(magic.JsonData, &data); err != nil {
		panic(err)
	}

	table := populateBishopAttackTables(data.Bishop)

	testCases := []struct {
		square        int
		blockers      uint64
		expectedMoves uint64
	}{
		{
			square:        sq.D4,
			blockers:      0,
			expectedMoves: 4693335752243822976,
		},
		{
			square:        sq.D4,
			blockers:      bitboard.SetBit(0, sq.B6),
			expectedMoves: 4693335752243822720,
		},
		{
			square:        sq.D4,
			blockers:      4692755210104356992,
			expectedMoves: 9029189822840832,
		},
		{
			square:        sq.A1,
			blockers:      bitboard.SetBit(0, sq.F6),
			expectedMoves: 567382630203392,
		},
		{
			square:        sq.E8,
			blockers:      bitboard.SetBits(0, sq.C6, sq.H5, sq.H4, sq.H3, sq.A2, sq.F8),
			expectedMoves: 2151950336,
		},
		{
			square:        sq.H1,
			blockers:      bitboard.SetBits(0, sq.D5, sq.C6, sq.B7, sq.A8, sq.B2, sq.B3, sq.C1, sq.C8),
			expectedMoves: 18049651735265280,
		},
	}

	for _, tt := range testCases {
		magicNum, err := strconv.ParseUint(data.Bishop.Magics[tt.square].Magic, 16, 64)
		if err != nil {
			panic(err)
		}

		mask, err := strconv.ParseUint(data.Bishop.Magics[tt.square].Mask, 16, 64)
		if err != nil {
			panic(err)
		}

		shift := data.Bishop.Magics[tt.square].Shift
		index := (tt.blockers & mask * magicNum) >> shift
		moves := table[tt.square][index]

		if moves != tt.expectedMoves {
			fmt.Println("Blockers:")
			bitboard.PrintBoard(tt.blockers)
			fmt.Println("")

			fmt.Println("Expected moves:")
			bitboard.PrintBoard(tt.expectedMoves)
			fmt.Println("")

			fmt.Println("Got moves:")
			bitboard.PrintBoard(moves)
			fmt.Println("")

			t.Error("For bishop on square", sq.Stringify(tt.square), "expected", tt.expectedMoves, "got", moves)
		}
	}
}
