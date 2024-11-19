package movegen

import (
	"fmt"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/internal/move"
	"github.com/samwestmoreland/chessengine/internal/piece"
	"github.com/samwestmoreland/chessengine/internal/position"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/internal/tables"
)

var lookupTables *tables.Lookup

// Initialise populates the lookup tables which are stored as a global variable in this package.
func Initialise() error {
	lookupTables = &tables.Lookup{
		Pawns:   [2][64]bb.Bitboard{},
		Knights: [64]bb.Bitboard{},
		Kings:   [64]bb.Bitboard{},
		Bishops: [64][]bb.Bitboard{},
		Rooks:   [64][]bb.Bitboard{},
	}

	err := tables.InitialiseLookupTables(lookupTables)
	if err != nil {
		return fmt.Errorf("failed to initialise lookup tables: %w", err)
	}

	return nil
}

func GetLegalMoves(pos *position.Position) []move.Move {
	var ret []move.Move

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

func SquareAttacked(pos *position.Position, square sq.Square, whiteAttacking bool) bool {
	if whiteAttacking {
		if lookupTables.Pawns[1][square]&pos.Occupancy[piece.Wp] != 0 {
			return true
		}

		if lookupTables.Kings[square]&pos.Occupancy[piece.Wk] != 0 {
			return true
		}

		if lookupTables.Knights[square]&pos.Occupancy[piece.Wn] != 0 {
			return true
		}

		bishopIndex := tables.GetBishopLookupIndex(square, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[piece.Wb] != 0 {
			return true
		}

		rookIndex := tables.GetRookLookupIndex(square, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])
		if lookupTables.Rooks[square][rookIndex]&pos.Occupancy[piece.Wr] != 0 {
			return true
		}

		// Lookup queen attacks using bishop and rook attacks
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[piece.Wq] != 0 ||
			lookupTables.Rooks[square][rookIndex]&pos.Occupancy[piece.Wq] != 0 {
			return true
		}
	} else {
		if lookupTables.Pawns[0][square]&pos.Occupancy[piece.Bp] != 0 {
			return true
		}

		if lookupTables.Kings[square]&pos.Occupancy[piece.Bk] != 0 {
			return true
		}

		if lookupTables.Knights[square]&pos.Occupancy[piece.Bn] != 0 {
			return true
		}

		bishopIndex := tables.GetBishopLookupIndex(square, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[piece.Bb] != 0 {
			return true
		}

		rookIndex := tables.GetRookLookupIndex(square, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])
		if lookupTables.Rooks[square][rookIndex]&pos.Occupancy[piece.Br] != 0 {
			return true
		}

		// Lookup queen attacks using bishop and rook attacks
		if lookupTables.Bishops[square][bishopIndex]&pos.Occupancy[piece.Bq] != 0 ||
			lookupTables.Rooks[square][rookIndex]&pos.Occupancy[piece.Bq] != 0 {
			return true
		}
	}

	return false
}

func getWhitePawnMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	var pawnOccupancy bb.Bitboard

	pawnOccupancy = pos.Occupancy[piece.Wp] // white specific

	for pawnOccupancy != 0 {
		source := bb.LSBIndex(pawnOccupancy)

		target := source - 8      // white specific
		doublePush := source - 16 // white specific

		// Pawn advances
		if sq.OnBoard(target) && !pos.IsOccupied(target) {
			// Check promotion
			if target < sq.A7 { // white specific
				promotionMoves := []move.Move{
					move.Encode(source, target, piece.Wp, piece.Wq, 0, 0, 0, 0), // white specific
					move.Encode(source, target, piece.Wp, piece.Wr, 0, 0, 0, 0), // white specific
					move.Encode(source, target, piece.Wp, piece.Wn, 0, 0, 0, 0), // white specific
					move.Encode(source, target, piece.Wp, piece.Wb, 0, 0, 0, 0), // white specific
				}

				ret = append(ret, promotionMoves...)
			} else {
				ret = append(ret,
					move.Encode(source, target, piece.Wp, piece.NoPiece, 0, 0, 0, 0), // white specific
				)

				// Double push
				if source >= sq.A2 && source <= sq.H2 && !pos.IsOccupied(doublePush) { // white specific
					ret = append(ret,
						move.Encode(source, doublePush, piece.Wp, piece.NoPiece, 0, 1, 0, 0), // white specific
					)
				}
			}
		}

		// Pawn captures
		attacks := lookupTables.Pawns[0][source] & pos.Occupancy[piece.Ba] // white specific

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if target < sq.A7 {
				promotionMoves := []move.Move{
					move.Encode(source, target, piece.Bp, piece.Bq, 1, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Br, 1, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Bn, 1, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Bb, 1, 0, 0, 0),
				}

				ret = append(ret, promotionMoves...)
			} else {
				ret = append(ret,
					move.Encode(source, target, piece.Bp, piece.NoPiece, 1, 0, 0, 0),
				)
			}

			attacks = bb.ClearBit(attacks, target)
		}

		if pos.EnPassantSquare != sq.NoSquare {
			enpassant := lookupTables.Pawns[0][source] & bb.SetBit(0, pos.EnPassantSquare)

			if enpassant != 0 {
				ret = append(ret,
					move.Encode(source, pos.EnPassantSquare, piece.Wp, piece.NoPiece, 0, 0, 1, 0),
				)
			}
		}

		pawnOccupancy = bb.ClearBit(pawnOccupancy, source)
	}

	return ret
}

func getBlackPawnMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	var pawnOccupancy bb.Bitboard

	pawnOccupancy = pos.Occupancy[piece.Bp]

	for pawnOccupancy != 0 {
		source := bb.LSBIndex(pawnOccupancy)

		target := source + 8
		doublePush := source + 16

		// Pawn advances
		if sq.OnBoard(target) && !pos.IsOccupied(target) {
			// Check promotion
			if target > sq.H2 {
				promotionMoves := []move.Move{
					move.Encode(source, target, piece.Bp, piece.Bq, 0, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Br, 0, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Bn, 0, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Bb, 0, 0, 0, 0),
				}

				ret = append(ret, promotionMoves...)
			} else {
				ret = append(ret,
					move.Encode(source, target, piece.Bp, piece.NoPiece, 0, 0, 0, 0),
				)

				// Doule push
				if source >= sq.A7 && source <= sq.H7 && !pos.IsOccupied(doublePush) {
					ret = append(ret,
						move.Encode(source, doublePush, piece.Bp, piece.NoPiece, 0, 1, 0, 0),
					)
				}
			}
		}

		// Pawn captures
		attacks := lookupTables.Pawns[1][source] & pos.Occupancy[piece.Wa]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			// Promotion
			if target > sq.H2 {
				promotionMoves := []move.Move{
					move.Encode(source, target, piece.Bp, piece.Bq, 1, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Br, 1, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Bn, 1, 0, 0, 0),
					move.Encode(source, target, piece.Bp, piece.Bb, 1, 0, 0, 0),
				}

				ret = append(ret, promotionMoves...)
			} else {
				ret = append(ret,
					move.Encode(source, target, piece.Bp, piece.NoPiece, 1, 0, 0, 0),
				)
			}

			attacks = bb.ClearBit(attacks, target)
		}

		if pos.EnPassantSquare != sq.NoSquare {
			enpassant := lookupTables.Pawns[1][source] & bb.SetBit(0, pos.EnPassantSquare)

			if enpassant != 0 {
				ret = append(ret,
					move.Encode(source, pos.EnPassantSquare, piece.Bp, piece.NoPiece, 0, 0, 1, 0),
				)
			}
		}

		pawnOccupancy = bb.ClearBit(pawnOccupancy, source)
	}

	return ret
}

func getWhiteKingCastlingMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	// King side castle
	if pos.CastlingRights&8 != 0 {
		// Check if pieces are in the way or if squares are attacked
		if !pos.IsOccupied(sq.F1) &&
			!pos.IsOccupied(sq.G1) &&
			!SquareAttacked(pos, sq.E1, false) &&
			!SquareAttacked(pos, sq.F1, false) &&
			!SquareAttacked(pos, sq.G1, false) {
			ret = append(ret, move.Encode(sq.E1, sq.G1, piece.Wk, piece.NoPiece, 0, 0, 0, 1))
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
			ret = append(ret, move.Encode(sq.E1, sq.C1, piece.Wk, piece.NoPiece, 0, 0, 0, 1))
		}
	}

	return ret
}

func getBlackKingCastlingMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	// King side castle
	if pos.CastlingRights&2 != 0 {
		// Check if pieces are in the way
		if !pos.IsOccupied(sq.F8) &&
			!pos.IsOccupied(sq.G8) &&
			!SquareAttacked(pos, sq.E8, true) &&
			!SquareAttacked(pos, sq.F8, true) &&
			!SquareAttacked(pos, sq.G8, true) {
			ret = append(ret, move.Encode(sq.E8, sq.G8, piece.Bk, piece.NoPiece, 0, 0, 0, 1))
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
			ret = append(ret, move.Encode(sq.E8, sq.C8, piece.Bk, piece.NoPiece, 0, 0, 0, 1))
		}
	}

	return ret
}

func getWhiteKnightMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	knights := pos.Occupancy[piece.Wn]

	for knights != 0 {
		source := bb.LSBIndex(knights)
		knights = bb.ClearBit(knights, source)

		attacks := lookupTables.Knights[source] &^ pos.Occupancy[piece.Wa]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Ba]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Wn, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Wn, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getBlackKnightMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	knights := pos.Occupancy[piece.Bn]

	for knights != 0 {
		source := bb.LSBIndex(knights)
		knights = bb.ClearBit(knights, source)

		attacks := lookupTables.Knights[source] &^ pos.Occupancy[piece.Ba]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Wa]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Bn, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Bn, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getWhiteBishopMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	bishops := pos.Occupancy[piece.Wb]

	for bishops != 0 {
		source := bb.LSBIndex(bishops)
		bishops = bb.ClearBit(bishops, source)

		index := tables.GetBishopLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		attacks := lookupTables.Bishops[source][index] &^ pos.Occupancy[piece.Wa]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Ba]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Wb, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Wb, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getBlackBishopMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	bishops := pos.Occupancy[piece.Bb]

	for bishops != 0 {
		source := bb.LSBIndex(bishops)
		bishops = bb.ClearBit(bishops, source)

		index := tables.GetBishopLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		attacks := lookupTables.Bishops[source][index] &^ pos.Occupancy[piece.Ba]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Wa]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Bb, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Bb, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getWhiteRookMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	rooks := pos.Occupancy[piece.Wr]

	for rooks != 0 {
		source := bb.LSBIndex(rooks)
		rooks = bb.ClearBit(rooks, source)

		index := tables.GetRookLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		attacks := lookupTables.Rooks[source][index] &^ pos.Occupancy[piece.Wa]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Ba]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Wr, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Wr, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getBlackRookMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	rooks := pos.Occupancy[piece.Br]

	for rooks != 0 {
		source := bb.LSBIndex(rooks)
		rooks = bb.ClearBit(rooks, source)

		index := tables.GetRookLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		attacks := lookupTables.Rooks[source][index] &^ pos.Occupancy[piece.Ba]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Wa]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Br, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Br, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getWhiteKingMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	king := pos.Occupancy[piece.Wk]

	for king != 0 {
		source := bb.LSBIndex(king)
		king = bb.ClearBit(king, source)

		attacks := lookupTables.Kings[source] &^ pos.Occupancy[piece.Wa]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Ba]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Wk, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Wk, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getBlackKingMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	king := pos.Occupancy[piece.Bk]

	for king != 0 {
		source := bb.LSBIndex(king)
		king = bb.ClearBit(king, source)

		attacks := lookupTables.Kings[source] &^ pos.Occupancy[piece.Ba]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if pos.Occupancy[piece.Wa]&(bb.SetBit(0, target)) != 0 {
				ret = append(ret, move.Encode(source, target, piece.Bk, piece.NoPiece, 1, 0, 0, 0))
			} else {
				ret = append(ret, move.Encode(source, target, piece.Bk, piece.NoPiece, 0, 0, 0, 0))
			}

			attacks = bb.ClearBit(attacks, target)
		}
	}

	return ret
}

func getWhiteQueenMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	queens := pos.Occupancy[piece.Wq]

	for queens != 0 {
		source := bb.LSBIndex(queens)
		queens = bb.ClearBit(queens, source)

		bishopTableIndex := tables.GetBishopLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		bishopAttacks := lookupTables.Bishops[source][bishopTableIndex] &^ pos.Occupancy[piece.Wa]

		rookTableIndex := tables.GetRookLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		rookAttacks := lookupTables.Rooks[source][rookTableIndex] &^ pos.Occupancy[piece.Wa]

		for bishopAttacks != 0 || rookAttacks != 0 {
			if bishopAttacks != 0 {
				target := bb.LSBIndex(bishopAttacks)

				if pos.Occupancy[piece.Ba]&(bb.SetBit(0, target)) != 0 {
					ret = append(ret, move.Encode(source, target, piece.Wq, piece.NoPiece, 1, 0, 0, 0))
				} else {
					ret = append(ret, move.Encode(source, target, piece.Wq, piece.NoPiece, 0, 0, 0, 0))
				}

				bishopAttacks = bb.ClearBit(bishopAttacks, target)
			}

			if rookAttacks != 0 {
				target := bb.LSBIndex(rookAttacks)

				if pos.Occupancy[piece.Ba]&(bb.SetBit(0, target)) != 0 {
					ret = append(ret, move.Encode(source, target, piece.Wq, piece.NoPiece, 1, 0, 0, 0))
				} else {
					ret = append(ret, move.Encode(source, target, piece.Wq, piece.NoPiece, 0, 0, 0, 0))
				}

				rookAttacks = bb.ClearBit(rookAttacks, target)
			}
		}
	}

	return ret
}

func getBlackQueenMoves(pos *position.Position) []move.Move {
	var ret []move.Move

	queens := pos.Occupancy[piece.Bq]

	for queens != 0 {
		source := bb.LSBIndex(queens)
		queens = bb.ClearBit(queens, source)

		bishopTableIndex := tables.GetBishopLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		bishopAttacks := lookupTables.Bishops[source][bishopTableIndex] &^ pos.Occupancy[piece.Ba]

		rookTableIndex := tables.GetRookLookupIndex(source, pos.Occupancy[piece.Wa]|pos.Occupancy[piece.Ba])

		rookAttacks := lookupTables.Rooks[source][rookTableIndex] &^ pos.Occupancy[piece.Ba]

		for bishopAttacks != 0 || rookAttacks != 0 {
			if bishopAttacks != 0 {
				target := bb.LSBIndex(bishopAttacks)

				if pos.Occupancy[piece.Wa]&(bb.SetBit(0, target)) != 0 {
					ret = append(ret, move.Encode(source, target, piece.Bq, piece.NoPiece, 1, 0, 0, 0))
				} else {
					ret = append(ret, move.Encode(source, target, piece.Bq, piece.NoPiece, 0, 0, 0, 0))
				}

				bishopAttacks = bb.ClearBit(bishopAttacks, target)
			}

			if rookAttacks != 0 {
				target := bb.LSBIndex(rookAttacks)

				if pos.Occupancy[piece.Wa]&(bb.SetBit(0, target)) != 0 {
					ret = append(ret, move.Encode(source, target, piece.Bq, piece.NoPiece, 1, 0, 0, 0))
				} else {
					ret = append(ret, move.Encode(source, target, piece.Bq, piece.NoPiece, 0, 0, 0, 0))
				}

				rookAttacks = bb.ClearBit(rookAttacks, target)
			}
		}
	}

	return ret
}
