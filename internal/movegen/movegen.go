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

// Initialise populates the lookup tables which are stored as a global variable in this package
func Initialise() error {
	lookupTables = &tables.Lookup{}

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

func SquareAttacked(pos *position.Position, square int, whiteAttacking bool) bool {
	if whiteAttacking {
		if lookupTables.Pawns[1][square]&pos.Occupancy[P] != 0 {
			return true
		}

		if lookupTables.Kings[square]&pos.Occupancy[K] != 0 {
			return true
		}

		if lookupTables.Knights[square]&pos.Occupancy[N] != 0 {
			return true
		}

		bishopIndex := tables.GetBishopLookupIndex(square, pos.Occupancy[A]|pos.Occupancy[a])
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[B] != 0 {
			return true
		}
	} else {
		if lookupTables.Pawns[0][square]&pos.Occupancy[p] != 0 {
			return true
		}

		if lookupTables.Kings[square]&pos.Occupancy[k] != 0 {
			return true
		}

		if lookupTables.Knights[square]&pos.Occupancy[n] != 0 {
			return true
		}

		bishopIndex := tables.GetBishopLookupIndex(square, pos.Occupancy[A]|pos.Occupancy[a])
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[b] != 0 {
			return true
		}
	}

	return false
}

func getWhitePawnMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	var pawnOccupancy bb.Bitboard

	pawnOccupancy = pos.Occupancy[P]

	for pawnOccupancy != 0 {
		source := bb.LSBIndex(pawnOccupancy)

		target := source - 8
		doublePush := source - 16

		// Pawn advances
		if sq.OnBoard(target) && !pos.IsOccupied(target) {
			// Check promotion
			if target < sq.A7 {
				ret = append(ret, getPawnPromotions(source, target)...)
			} else {
				ret = append(ret,
					position.Move{From: source, To: target},
				)

				// Double push
				if source >= sq.A2 && source <= sq.H2 && !pos.IsOccupied(doublePush) {
					ret = append(ret,
						position.Move{From: source, To: doublePush},
					)
				}
			}
		}

		// Pawn captures
		attacks := lookupTables.Pawns[0][source] & pos.Occupancy[a]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if target < sq.A7 {
				ret = append(ret, getPawnPromotions(source, target)...)
			} else {
				ret = append(ret,
					position.Move{From: source, To: target},
				)
			}

			attacks = bb.ClearBit(attacks, target)
		}

		if pos.EnPassantSquare != sq.NoSquare {
			enpassant := lookupTables.Pawns[0][source] & bb.SetBit(0, int(pos.EnPassantSquare))

			if enpassant != 0 {
				ret = append(ret,
					position.Move{From: source, To: int(pos.EnPassantSquare)},
				)
			}
		}

		pawnOccupancy = bb.ClearBit(pawnOccupancy, source)
	}

	return ret
}

func getBlackPawnMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	var pawnOccupancy bb.Bitboard

	pawnOccupancy = pos.Occupancy[p]

	for pawnOccupancy != 0 {
		source := bb.LSBIndex(pawnOccupancy)

		target := source + 8
		doublePush := source + 16

		// Pawn advances
		if sq.OnBoard(target) && !pos.IsOccupied(target) {
			// Check promotion
			if target > sq.H2 {
				ret = append(ret, getPawnPromotions(source, target)...)
			} else {
				ret = append(ret,
					position.Move{From: source, To: target},
				)

				// Doule push
				if source >= sq.A7 && source <= sq.H7 && !pos.IsOccupied(doublePush) {
					ret = append(ret,
						position.Move{From: source, To: doublePush},
					)
				}
			}
		}

		// Pawn captures
		attacks := lookupTables.Pawns[1][source] & pos.Occupancy[A]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if target > sq.H2 {
				ret = append(ret, getPawnPromotions(source, target)...)
			} else {
				ret = append(ret,
					position.Move{From: source, To: target},
				)
			}

			attacks = bb.ClearBit(attacks, target)
		}

		if pos.EnPassantSquare != sq.NoSquare {
			enpassant := lookupTables.Pawns[1][source] & bb.SetBit(0, int(pos.EnPassantSquare))

			if enpassant != 0 {
				ret = append(ret,
					position.Move{From: source, To: int(pos.EnPassantSquare)},
				)
			}
		}

		pawnOccupancy = bb.ClearBit(pawnOccupancy, source)
	}

	return ret
}

func getPawnPromotions(source, target int) []position.Move {
	return []position.Move{
		{From: source, To: target, PromotionPiece: "q"},
		{From: source, To: target, PromotionPiece: "r"},
		{From: source, To: target, PromotionPiece: "b"},
		{From: source, To: target, PromotionPiece: "n"},
	}
}

func getWhiteKingCastlingMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	// King side castle
	if pos.CastlingRights&8 != 0 {
		// Check if pieces are in the way
		if !pos.IsOccupied(sq.F1) && !pos.IsOccupied(sq.G1) {
			ret = append(ret, position.Move{From: sq.E1, To: sq.G1})
		}
	}

	// Queen side castle
	if pos.CastlingRights&4 != 0 {
		// Check if pieces are in the way
		if !pos.IsOccupied(sq.D1) && !pos.IsOccupied(sq.C1) && !pos.IsOccupied(sq.B1) {
			ret = append(ret, position.Move{From: sq.E1, To: sq.C1})
		}
	}

	return ret
}

// func getBishopAttacks(square int) bb.Bitboard {
// 	return lookupTables.Bishops[square] & (pos.Occupancy[A] | pos.Occupancy[a])
// }
