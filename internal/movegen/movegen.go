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
		ret = append(ret, getPawnMoves(pos, piece.White)...)
		ret = append(ret, getWhiteKingCastlingMoves(pos)...)
		ret = append(ret, getWhiteKnightMoves(pos)...)
		ret = append(ret, getWhiteBishopMoves(pos)...)
		ret = append(ret, getWhiteRookMoves(pos)...)
		ret = append(ret, getWhiteKingMoves(pos)...)
		ret = append(ret, getQueenMoves(pos, piece.White)...)
	} else {
		ret = append(ret, getPawnMoves(pos, piece.Black)...)
		ret = append(ret, getBlackKingCastlingMoves(pos)...)
		ret = append(ret, getBlackKnightMoves(pos)...)
		ret = append(ret, getBlackBishopMoves(pos)...)
		ret = append(ret, getBlackRookMoves(pos)...)
		ret = append(ret, getBlackKingMoves(pos)...)
		ret = append(ret, getQueenMoves(pos, piece.Black)...)
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

func getPawnMoves(pos *position.Position, colour piece.Colour) []move.Move {
	var ret []move.Move

	var pawnPiece, queenPiece, rookPiece, bishopPiece, knightPiece piece.Piece

	var enemyPieces piece.Piece

	var penultimateRank, startRank int

	if colour == piece.White {
		pawnPiece = piece.Wp
		queenPiece = piece.Wq
		rookPiece = piece.Wr
		bishopPiece = piece.Wb
		knightPiece = piece.Wn
		enemyPieces = piece.Ba
		penultimateRank = 7
		startRank = 2
	} else {
		pawnPiece = piece.Bp
		queenPiece = piece.Bq
		rookPiece = piece.Br
		bishopPiece = piece.Bb
		knightPiece = piece.Bn
		enemyPieces = piece.Wa
		penultimateRank = 2
		startRank = 7
	}

	pawns := pos.Occupancy[pawnPiece]

	for pawns != 0 {
		source := bb.LSBIndex(pawns)

		var target, doublePush sq.Square
		if colour == piece.White {
			target = source - 8
			doublePush = source - 16
		} else {
			target = source + 8
			doublePush = source + 16
		}

		// Pawn advances
		if sq.OnBoard(target) && !pos.IsOccupied(target) {
			// Check promotion
			if bb.IsNthRank(penultimateRank, source) {
				promotionMoves := []move.Move{
					move.Encode(source, target, pawnPiece, queenPiece, 0, 0, 0, 0),
					move.Encode(source, target, pawnPiece, rookPiece, 0, 0, 0, 0),
					move.Encode(source, target, pawnPiece, bishopPiece, 0, 0, 0, 0),
					move.Encode(source, target, pawnPiece, knightPiece, 0, 0, 0, 0),
				}

				ret = append(ret, promotionMoves...)
			} else {
				ret = append(ret,
					move.Encode(source, target, pawnPiece, piece.NoPiece, 0, 0, 0, 0),
				)

				// Double push
				if bb.IsNthRank(startRank, source) && !pos.IsOccupied(doublePush) {
					ret = append(ret,
						move.Encode(source, doublePush, pawnPiece, piece.NoPiece, 0, 1, 0, 0),
					)
				}
			}
		}

		// Pawn captures
		attacks := lookupTables.Pawns[colour][source] & pos.Occupancy[enemyPieces]

		for attacks != 0 {
			target := bb.LSBIndex(attacks)

			if bb.IsNthRank(penultimateRank, source) {
				promotionMoves := []move.Move{
					move.Encode(source, target, pawnPiece, queenPiece, 1, 0, 0, 0),
					move.Encode(source, target, pawnPiece, rookPiece, 1, 0, 0, 0),
					move.Encode(source, target, pawnPiece, bishopPiece, 1, 0, 0, 0),
					move.Encode(source, target, pawnPiece, knightPiece, 1, 0, 0, 0),
				}

				ret = append(ret, promotionMoves...)
			} else {
				ret = append(ret,
					move.Encode(source, target, pawnPiece, piece.NoPiece, 1, 0, 0, 0),
				)
			}

			attacks = bb.ClearBit(attacks, target)
		}

		if pos.EnPassantSquare != sq.NoSquare {
			enpassant := lookupTables.Pawns[colour][source] & bb.SetBit(0, pos.EnPassantSquare)

			if enpassant != 0 {
				ret = append(ret,
					move.Encode(source, pos.EnPassantSquare, pawnPiece, piece.NoPiece, 0, 0, 1, 0),
				)
			}
		}

		pawns = bb.ClearBit(pawns, source)
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

func getQueenMoves(pos *position.Position, colour piece.Colour) []move.Move {
	var ret []move.Move

	var queenPiece piece.Piece

	var ownOccupancy, enemyOccupancy piece.Piece

	if colour == piece.White {
		queenPiece = piece.Wq
		ownOccupancy = piece.Wa
		enemyOccupancy = piece.Ba
	} else {
		queenPiece = piece.Bq
		ownOccupancy = piece.Ba
		enemyOccupancy = piece.Wa
	}

	queens := pos.Occupancy[queenPiece]

	for queens != 0 {
		source := bb.LSBIndex(queens)
		queens = bb.ClearBit(queens, source)

		allPieces := pos.Occupancy[ownOccupancy] | pos.Occupancy[enemyOccupancy]
		bishopTableIndex := tables.GetBishopLookupIndex(source, allPieces)
		rookTableIndex := tables.GetRookLookupIndex(source, allPieces)

		bishopAttacks := lookupTables.Bishops[source][bishopTableIndex] &^ pos.Occupancy[ownOccupancy]
		rookAttacks := lookupTables.Rooks[source][rookTableIndex] &^ pos.Occupancy[ownOccupancy]

		// Helper function to avoid duplication in move generation
		addMoves := func(attacks bb.Bitboard) {
			for attacks != 0 {
				target := bb.LSBIndex(attacks)

				var captureFlag uint32
				if pos.Occupancy[enemyOccupancy]&(bb.SetBit(0, target)) != 0 {
					captureFlag = 1
				}

				ret = append(ret, move.Encode(source, target, queenPiece, piece.NoPiece, captureFlag, 0, 0, 0))
				attacks = bb.ClearBit(attacks, target)
			}
		}

		addMoves(bishopAttacks)
		addMoves(rookAttacks)
	}

	return ret
}
