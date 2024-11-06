package tables

import (
	"encoding/json"

	"github.com/samwestmoreland/chessengine/magic"
)

type Lookup struct {
	Pawns   [2][64]uint64
	Knights [64]uint64
	Kings   [64]uint64
	Bishops [64][]uint64
	Rooks   [64][]uint64
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
