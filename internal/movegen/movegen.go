package movegen

import (
	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/internal/position"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/internal/tables"
)

var lookupTables *tables.Lookup

const (
	P = iota
	N
	B
	R
	Q
	K
	p
	n
	b
	r
	q
	k
	A // All white
	a // All black
)

func Initialise() error {
	lookupTables := &tables.Lookup{}

	return tables.InitialiseLookupTables(lookupTables)
}

func GetLegalMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	if pos.WhiteToMove {
		ret = append(ret, getWhitePawnMoves(pos)...)
	} else {
		ret = append(ret, getBlackPawnMoves(pos)...)
	}

	return ret
}

func getWhitePawnMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	var pawnOccupancy bb.Bitboard

	pawnOccupancy = pos.Occupancy[P]

	for pawnOccupancy != 0 {
		source := bb.LSBIndex(pawnOccupancy)
		pawnOccupancy = bb.ClearBit(pawnOccupancy, source)

		target := source - 8
		doublePush := source - 16

		// Pawn advances
		if sq.OnBoard(target) && !pos.IsOccupied(target) {
			// Check promotion
			if target < sq.A7 {
				ret = append(ret,
					position.Move{From: source, To: target, PromotionPiece: "q"},
					position.Move{From: source, To: target, PromotionPiece: "r"},
					position.Move{From: source, To: target, PromotionPiece: "b"},
					position.Move{From: source, To: target, PromotionPiece: "n"},
				)
			} else {
				ret = append(ret,
					position.Move{From: source, To: target},
				)

				if source >= sq.A2 && source <= sq.H2 && !pos.IsOccupied(doublePush) {
					ret = append(ret,
						position.Move{From: source, To: doublePush},
					)
				}
			}
		}
	}

	return ret
}

func getBlackPawnMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	var pawnOccupancy bb.Bitboard

	pawnOccupancy = pos.Occupancy[p]

	for pawnOccupancy != 0 {
		source := bb.LSBIndex(pawnOccupancy)
		pawnOccupancy = bb.ClearBit(pawnOccupancy, source)

		target := source + 8
		doublePush := source + 16

		// Pawn advances
		if sq.OnBoard(target) && !pos.IsOccupied(target) {
			// Check promotion
			if target < sq.A7 {
				ret = append(ret,
					position.Move{From: source, To: target, PromotionPiece: "q"},
					position.Move{From: source, To: target, PromotionPiece: "r"},
					position.Move{From: source, To: target, PromotionPiece: "b"},
					position.Move{From: source, To: target, PromotionPiece: "n"},
				)
			} else {
				ret = append(ret,
					position.Move{From: source, To: target},
				)

				if source >= sq.A2 && source <= sq.H2 && !pos.IsOccupied(doublePush) {
					ret = append(ret,
						position.Move{From: source, To: doublePush},
					)
				}
			}
		}
	}

	return ret
}
