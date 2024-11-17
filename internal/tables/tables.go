package tables

import (
	"encoding/json"
	"fmt"
	"strconv"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
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
	if err := json.Unmarshal(magic.JSONData, &data); err != nil {
		return fmt.Errorf("failed to unmarshal magic data: %w", err)
	}

	table.Pawns = populatePawnAttackTables()
	table.Knights = populateKnightAttackTables()
	table.Kings = populateKingAttackTables()
	table.Bishops = populateBishopAttackTables(data.Bishop)
	table.Rooks = populateRookAttackTables(data.Rook)

	return nil
}

func GetBishopLookupIndex(square sq.Square, blockers bb.Bitboard) bb.Bitboard {
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

func GetRookLookupIndex(square sq.Square, blockers bb.Bitboard) bb.Bitboard {
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
