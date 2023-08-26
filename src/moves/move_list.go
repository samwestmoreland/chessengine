package moves

import (
	"github.com/samwestmoreland/chessengine/src/piece"
)

// MoveList is a list of moves.
// TODO: Make this a set. And we probably want to store hash
// values of the moves to make comparisons faster.
type MoveList struct {
	moves []Move
}

func (ml *MoveList) AddMove(m Move) {
	ml.moves = append(ml.moves, m)
}

func (ml *MoveList) GetMoves() []Move {
	return ml.moves
}

func (ml *MoveList) GetMovesForPieceType(pieceType piece.Type) []Move {
	var moves []Move
	for _, move := range ml.moves {
		if move.GetPieceType() == pieceType {
			moves = append(moves, move)
		}
	}
	return moves
}

// Equals checks if two MoveLists are equal, without worrying about order
func (ml *MoveList) Equals(other MoveList) bool {
	if len(ml.moves) != len(other.moves) {
		return false
	}
	for _, move := range ml.moves {
		if !other.Contains(move) {
			return false
		}
	}
	return true
}

// Contains checks if a MoveList contains a given move
func (ml *MoveList) Contains(move Move) bool {
	for _, m := range ml.moves {
		if m.Equals(move) {
			return true
		}
	}
	return false
}
