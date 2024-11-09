package tables

import (
	"encoding/json"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/magic"
)

type Lookup struct {
	Pawns   [2][64]bb.Bitboard
	Knights [64]bb.Bitboard
	Kings   [64]bb.Bitboard
	Bishops [64][]bb.Bitboard
	Rooks   [64][]bb.Bitboard
}

func InitialiseLookupTables(table *Lookup) error {
	var data magic.Data
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
