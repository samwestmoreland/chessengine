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
		ret = append(ret, getWhiteKingCastlingMoves(pos)...)
		ret = append(ret, getWhiteKnightMoves(pos)...)
		ret = append(ret, getWhiteBishopMoves(pos)...)
		ret = append(ret, getWhiteRookMoves(pos)...)
		ret = append(ret, getWhiteKingMoves(pos)...)
		ret = append(ret, getWhiteQueenMoves(pos)...)
	} else {
		ret = append(ret, getBlackPawnMoves(pos)...)
		ret = append(ret, getBlackKingCastlingMoves(pos)...)
		ret = append(ret, getBlackKnightMoves(pos)...)
		ret = append(ret, getBlackBishopMoves(pos)...)
		ret = append(ret, getBlackRookMoves(pos)...)
		ret = append(ret, getBlackKingMoves(pos)...)
		ret = append(ret, getBlackQueenMoves(pos)...)
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

		rookIndex := tables.GetRookLookupIndex(square, pos.Occupancy[A]|pos.Occupancy[a])
		if lookupTables.Rooks[square][rookIndex]&pos.Occupancy[R] != 0 {
			return true
		}

		// Lookup queen attacks using bishop and rook attacks
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[Q] != 0 ||
			lookupTables.Rooks[square][rookIndex]&pos.Occupancy[Q] != 0 {
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

		rookIndex := tables.GetRookLookupIndex(square, pos.Occupancy[A]|pos.Occupancy[a])
		if lookupTables.Rooks[square][rookIndex]&pos.Occupancy[r] != 0 {
			return true
		}

		// Lookup queen attacks using bishop and rook attacks
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[q] != 0 ||
			lookupTables.Rooks[square][rookIndex]&pos.Occupancy[q] != 0 {
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
		// Check if pieces are in the way or if squares are attacked
		if !pos.IsOccupied(sq.F1) &&
			!pos.IsOccupied(sq.G1) &&
			!SquareAttacked(pos, sq.E1, false) &&
			!SquareAttacked(pos, sq.F1, false) &&
			!SquareAttacked(pos, sq.G1, false) {
			ret = append(ret, position.Move{From: sq.E1, To: sq.G1})
		}
	}

	// Queen side castle
	if pos.CastlingRights&4 != 0 {
		// Check if pieces are in the way
		if !pos.IsOccupied(sq.D1) &&
			!pos.IsOccupied(sq.C1) &&
			!pos.IsOccupied(sq.B1) &&
			!SquareAttacked(pos, sq.E1, false) &&
			!SquareAttacked(pos, sq.D1, false) &&
			!SquareAttacked(pos, sq.C1, false) {
			ret = append(ret, position.Move{From: sq.E1, To: sq.C1})
		}
	}

	return ret
}

func getBlackKingCastlingMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	// King side castle
	if pos.CastlingRights&2 != 0 {
		// Check if pieces are in the way
		if !pos.IsOccupied(sq.F8) &&
			!pos.IsOccupied(sq.G8) &&
			!SquareAttacked(pos, sq.E8, true) &&
			!SquareAttacked(pos, sq.F8, true) &&
			!SquareAttacked(pos, sq.G8, true) {
			ret = append(ret, position.Move{From: sq.E8, To: sq.G8})
		}
	}

	// Queen side castle
	if pos.CastlingRights&1 != 0 {
		// Check if pieces are in the way
		if !pos.IsOccupied(sq.D8) &&
			!pos.IsOccupied(sq.C8) &&
			!pos.IsOccupied(sq.B8) &&
			!SquareAttacked(pos, sq.E8, true) &&
			!SquareAttacked(pos, sq.D8, true) &&
			!SquareAttacked(pos, sq.C8, true) {
			ret = append(ret, position.Move{From: sq.E8, To: sq.C8})
		}
	}

	return ret
}

func getWhiteKnightMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	knights := pos.Occupancy[N]

	for knights != 0 {
		source := bb.LSBIndex(knights)
		knights = bb.ClearBit(knights, source)

		attacks := lookupTables.Knights[source] &^ pos.Occupancy[A]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getBlackKnightMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	knights := pos.Occupancy[n]

	for knights != 0 {
		source := bb.LSBIndex(knights)
		knights = bb.ClearBit(knights, source)

		attacks := lookupTables.Knights[source] &^ pos.Occupancy[a]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getWhiteBishopMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	bishops := pos.Occupancy[B]

	for bishops != 0 {
		source := bb.LSBIndex(bishops)
		bishops = bb.ClearBit(bishops, source)

		index := tables.GetBishopLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		attacks := lookupTables.Bishops[source][index] &^ pos.Occupancy[A]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getBlackBishopMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	bishops := pos.Occupancy[b]

	for bishops != 0 {
		source := bb.LSBIndex(bishops)
		bishops = bb.ClearBit(bishops, source)

		index := tables.GetBishopLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		attacks := lookupTables.Bishops[source][index] &^ pos.Occupancy[a]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getWhiteRookMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	rooks := pos.Occupancy[R]

	for rooks != 0 {
		source := bb.LSBIndex(rooks)
		rooks = bb.ClearBit(rooks, source)

		index := tables.GetRookLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		attacks := lookupTables.Rooks[source][index] &^ pos.Occupancy[A]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getBlackRookMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	rooks := pos.Occupancy[r]

	for rooks != 0 {
		source := bb.LSBIndex(rooks)
		rooks = bb.ClearBit(rooks, source)

		index := tables.GetRookLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		attacks := lookupTables.Rooks[source][index] &^ pos.Occupancy[a]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getWhiteKingMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	king := pos.Occupancy[K]

	for king != 0 {
		source := bb.LSBIndex(king)
		king = bb.ClearBit(king, source)

		attacks := lookupTables.Kings[source] &^ pos.Occupancy[A]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getBlackKingMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	king := pos.Occupancy[k]

	for king != 0 {
		source := bb.LSBIndex(king)
		king = bb.ClearBit(king, source)

		attacks := lookupTables.Kings[source] &^ pos.Occupancy[a]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)
			attacks = bb.ClearBit(attacks, target)

			ret = append(ret, position.Move{From: source, To: target})
		}
	}

	return ret
}

func getWhiteQueenMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	queens := pos.Occupancy[Q]

	for queens != 0 {
		source := bb.LSBIndex(queens)
		queens = bb.ClearBit(queens, source)

		bishopTableIndex := tables.GetBishopLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		bishopAttacks := lookupTables.Bishops[source][bishopTableIndex] &^ pos.Occupancy[A]

		rookTableIndex := tables.GetRookLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		rookAttacks := lookupTables.Rooks[source][rookTableIndex] &^ pos.Occupancy[A]

		for bishopAttacks != 0 || rookAttacks != 0 {
			if bishopAttacks != 0 {
				target := bb.LSBIndex(bishopAttacks)
				bishopAttacks = bb.ClearBit(bishopAttacks, target)

				ret = append(ret, position.Move{From: source, To: target})
			}

			if rookAttacks != 0 {
				target := bb.LSBIndex(rookAttacks)
				rookAttacks = bb.ClearBit(rookAttacks, target)

				ret = append(ret, position.Move{From: source, To: target})
			}
		}
	}

	return ret
}

func getBlackQueenMoves(pos *position.Position) []position.Move {
	var ret []position.Move

	queens := pos.Occupancy[q]

	for queens != 0 {
		source := bb.LSBIndex(queens)
		queens = bb.ClearBit(queens, source)

		bishopTableIndex := tables.GetBishopLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		bishopAttacks := lookupTables.Bishops[source][bishopTableIndex] &^ pos.Occupancy[a]

		rookTableIndex := tables.GetRookLookupIndex(source, pos.Occupancy[A]|pos.Occupancy[a])

		rookAttacks := lookupTables.Rooks[source][rookTableIndex] &^ pos.Occupancy[a]

		for bishopAttacks != 0 || rookAttacks != 0 {
			if bishopAttacks != 0 {
				target := bb.LSBIndex(bishopAttacks)
				bishopAttacks = bb.ClearBit(bishopAttacks, target)

				ret = append(ret, position.Move{From: source, To: target})
			}

			if rookAttacks != 0 {
				target := bb.LSBIndex(rookAttacks)
				rookAttacks = bb.ClearBit(rookAttacks, target)

				ret = append(ret, position.Move{From: source, To: target})
			}
		}
	}

	return ret
}
