package eval

import (
	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/piece"
	"github.com/samwestmoreland/chessengine/src/position"
)

type ShannonEvaluator struct{}

func (e ShannonEvaluator) Evaluate(pos *position.Position) float64 {
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

	whitePawns := pos.GetPiecesForColour(piece.PawnType, board.White)

	for _, file := range board.Files {
		found := false
		for _, pawn := range whitePawns {
			if !found && pawn.File == file {
				found = true
			} else if found && pawn.File == file {


	score := float64(
		queenDifference*9 +
			rookDifference*5 +
			bishopDifference*3 +
			knightDifference*3 +
			pawnDifference,
	)

	return score
}
