package eval

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/piece"
	"github.com/samwestmoreland/chessengine/src/position"
)

// ShannonEvaluator is an implementation of the Evaluator interface. It uses the Shannon evaluation function.
// Currently we ignore the backward pawn penalty because I haven't figured out how to calculate it.
type ShannonEvaluator struct{}

func NewShannonEvaluator() ShannonEvaluator {
	return ShannonEvaluator{}
}

func (e ShannonEvaluator) Evaluate(pos *position.Position) float64 {
	materialScore := e.getMaterialScore(pos)
	pawnScore := e.getPawnScore(pos)
	mobilityScore := e.getMobilityScore(pos)

	score := float64(materialScore + pawnScore + mobilityScore)

	return score
}

func (e ShannonEvaluator) getMaterialScore(pos *position.Position) int {
	queenDifference := pos.GetNumPiecesForColour(piece.QueenType, board.White) -
		pos.GetNumPiecesForColour(piece.QueenType, board.Black)

	rookDifference := pos.GetNumPiecesForColour(piece.RookType, board.White) -
		pos.GetNumPiecesForColour(piece.RookType, board.Black)

	bishopDifference := pos.GetNumPiecesForColour(piece.BishopType, board.White) -
		pos.GetNumPiecesForColour(piece.BishopType, board.Black)

	knightDifference := pos.GetNumPiecesForColour(piece.KnightType, board.White) -
		pos.GetNumPiecesForColour(piece.KnightType, board.Black)

	pawnDifference := pos.GetNumPiecesForColour(piece.PawnType, board.White) -
		pos.GetNumPiecesForColour(piece.PawnType, board.Black)

	return queenDifference*9 + rookDifference*5 + bishopDifference*3 + knightDifference*3 + pawnDifference
}

func (e ShannonEvaluator) getPawnScore(pos *position.Position) int {
	var doubledPawnsWhite, isolatedPawnsWhite int

	whitePawns := pos.GetPiecesForColour(piece.PawnType, board.White)
	for _, pi := range whitePawns {
		if p, ok := pi.(position.Pawn); ok {
			if p.IsDoubled(pos) {
				doubledPawnsWhite++
			}

			if p.IsIsolated(pos) {
				isolatedPawnsWhite++
			}
		}
	}

	var doubledPawnsBlack, isolatedPawnsBlack int

	blackPawns := pos.GetPiecesForColour(piece.PawnType, board.Black)
	for _, pi := range blackPawns {
		if p, ok := pi.(position.Pawn); ok {
			if p.IsDoubled(pos) {
				doubledPawnsBlack++
			}

			if p.IsIsolated(pos) {
				isolatedPawnsBlack++
			}
		}
	}

	return doubledPawnsWhite - doubledPawnsBlack + isolatedPawnsWhite - isolatedPawnsBlack
}

func (e ShannonEvaluator) getMobilityScore(pos *position.Position) int {
	allWhiteMoves, err := pos.GetAllMovesConcurrent(board.White)
	if err != nil {
		panic(err)
	}

	allBlackMoves, err := pos.GetAllMovesConcurrent(board.Black)
	if err != nil {
		panic(err)
	}

	whiteMobility := allWhiteMoves.Len()
	blackMobility := allBlackMoves.Len()

	return whiteMobility - blackMobility
}
