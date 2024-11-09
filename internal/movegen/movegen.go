package movegen

import (
	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/internal/position"
	"github.com/samwestmoreland/chessengine/internal/tables"
)

var lookupTables *tables.Lookup

func Initialise() error {
	lookupTables := &tables.Lookup{}

	return tables.InitialiseLookupTables(lookupTables)
}

func GetLegalMoves(pos *position.Position) []bb.Bitboard {
	return nil
}
