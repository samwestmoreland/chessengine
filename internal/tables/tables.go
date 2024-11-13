package tables

import (
	"encoding/json"
	"strconv"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/magic"
)

var data magic.Data

type Lookup struct {
	Pawns   [2][64]bb.Bitboard
	Knights [64]bb.Bitboard
	Kings   [64]bb.Bitboard
	Bishops [64][]bb.Bitboard
	Rooks   [64][]bb.Bitboard
}

func InitialiseLookupTables(table *Lookup) error {
	data = data
	if err := json.Unmarshal(magic.JsonData, &data); err != nil {
		return err
	}

	table.Pawns = populatePawnAttackTables()
	table.Knights = populateKnightAttackTables()
	table.Kings = populateKingAttackTables()
	table.Bishops = populateBishopAttackTables(data.Bishop)
	table.Rooks = populateRookAttackTables(data.Rook)

	return nil
}

func GetBishopLookupIndex(square int, blockers bb.Bitboard) bb.Bitboard {
	magicNum, err := strconv.ParseUint(data.Bishop.Magics[square].Magic, 16, 64)
	if err != nil {
		panic(err)
	}

	mask, err := strconv.ParseUint(data.Bishop.Magics[square].Mask, 16, 64)
	if err != nil {
		panic(err)
	}

	shift := data.Bishop.Magics[square].Shift

	return bb.Bitboard((uint64(blockers) & mask * magicNum) >> shift)
}

func GetRookLookupIndex(square int, blockers bb.Bitboard) bb.Bitboard {
	magicNum, err := strconv.ParseUint(data.Rook.Magics[square].Magic, 16, 64)
	if err != nil {
		panic(err)
	}

	mask, err := strconv.ParseUint(data.Rook.Magics[square].Mask, 16, 64)
	if err != nil {
		panic(err)
	}

	shift := data.Rook.Magics[square].Shift

	return bb.Bitboard((uint64(blockers) & mask * magicNum) >> shift)
}
