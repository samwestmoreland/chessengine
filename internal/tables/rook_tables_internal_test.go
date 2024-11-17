package tables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/magic"
)

func TestMaskRookAttacks(t *testing.T) {
	var tests = map[sq.Square]uint64{
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
			var buf bytes.Buffer

			buf.WriteString(fmt.Sprintf("Square %s\n", sq.Stringify(square)))

			buf.WriteString("Got")
			bb.PrintBoard(actual, &buf)

			buf.WriteString("Expected")
			bb.PrintBoard(bb.Bitboard(expected), &buf)

			buf.WriteString("\n")

			t.Error(buf.String())

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
		var buf bytes.Buffer

		buf.WriteString("Blockers:\n")
		bb.PrintBoard(blockers, &buf)

		buf.WriteString("\nRook attacks:\n")
		bb.PrintBoard(RookAttacksOnTheFly(sq.D4, blockers), &buf)

		t.Error(buf.String())
		t.Error("Expected 9028156000256, got ", rookAttacks)
	}
}

func TestLookupTableGivesCorrectMovesForRook(t *testing.T) {
	var data magic.Data
	if err := json.Unmarshal(magic.JSONData, &data); err != nil {
		panic(err)
	}

	table := populateRookAttackTables(data.Rook)

	testCases := []struct {
		square        sq.Square
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
			var buf bytes.Buffer

			buf.WriteString("Blockers:")
			bb.PrintBoard(tt.blockers, &buf)
			buf.WriteString("\n")

			buf.WriteString("Expected moves:")
			bb.PrintBoard(bb.Bitboard(tt.expectedMoves), &buf)
			buf.WriteString("\n")

			buf.WriteString("Got moves:")
			bb.PrintBoard(moves, &buf)
			buf.WriteString("\n")

			t.Error(buf.String())

			t.Error("For rook on square", sq.Stringify(tt.square), "expected", tt.expectedMoves, "got", moves)
		}
	}
}
