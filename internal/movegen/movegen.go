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
	targetBit := bb.Bitboard(1) << square

	// Check if attacked by pawn
	if whiteAttacking {
		whitePawns := pos.Occupancy[P]

		for whitePawns != 0 {
			pawnSquare := bb.LSBIndex(whitePawns)
			if lookupTables.Pawns[0][pawnSquare]&targetBit != 0 {
				return true
			}

			whitePawns = bb.ClearBit(whitePawns, pawnSquare)
		}

		// Check if attacked by knight
		whiteKnights := pos.Occupancy[N]

		for whiteKnights != 0 {
			knightSquare := bb.LSBIndex(whiteKnights)
			if lookupTables.Knights[knightSquare]&targetBit != 0 {
				return true
			}

			whiteKnights = bb.ClearBit(whiteKnights, knightSquare)
		}
	}

	blackPawns := pos.Occupancy[p]

	for blackPawns != 0 {
		pawnSquare := bb.LSBIndex(blackPawns)
		if lookupTables.Pawns[1][pawnSquare]&targetBit != 0 {
			return true
		}

		blackPawns = bb.ClearBit(blackPawns, pawnSquare)
	}

	// Check if attacked by knight
	blackKnights := pos.Occupancy[n]

	for blackKnights != 0 {
		knightSquare := bb.LSBIndex(blackKnights)
		if lookupTables.Knights[knightSquare]&targetBit != 0 {
			return true
		}

		blackKnights = bb.ClearBit(blackKnights, knightSquare)
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
