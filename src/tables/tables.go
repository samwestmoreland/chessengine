package tables

import (
	"encoding/json"

	"github.com/samwestmoreland/chessengine/magic"
)

type Lookup struct {
	Pawns   [2][64]uint64
	Knights [64]uint64
	Kings   [64]uint64
	Bishops [64][512]uint64
	Rooks   [64][4096]uint64
}

func InitialiseLookupTables() (*Lookup, error) {
	// Load magic json data
	var data magic.Data
	if err := json.Unmarshal(magic.JsonData, &data); err != nil {
		return nil, err
	}

	return &Lookup{
		Pawns:   populatePawnAttackTables(),
		Knights: populateKnightAttackTables(),
		Kings:   populateKingAttackTables(),
		Bishops: populateBishopAttackTables(data.Bishop),
		Rooks:   populateRookAttackTables(data.Rook),
	}, nil
}
