package moves

import (
	"strings"

	"github.com/samwestmoreland/chessengine/internal/piece"
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

func (ml *MoveList) AddMoves(moves []Move) {
	ml.moves = append(ml.moves, moves...)
}

func (ml *MoveList) AddMoveList(moveList MoveList) {
	ml.moves = append(ml.moves, moveList.moves...)
}

func (ml *MoveList) Len() int {
	return len(ml.moves)
}

func (ml *MoveList) GetMoves() []Move {
	return ml.moves
}

func (ml *MoveList) GetMovesForPieceType(pieceType piece.Type) MoveList {
	var ret MoveList

	for _, move := range ml.moves {
		if move.GetPieceType() == pieceType {
			ret.AddMove(move)
		}
	}

	return ret
}

// Equals checks if two MoveLists are equal, without worrying about order.
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

// Contains checks if a MoveList contains a given move.
func (ml *MoveList) Contains(move Move) bool {
	for _, m := range ml.moves {
		if m.Equals(move) {
			return true
		}
	}

	return false
}

// String returns the MoveList in an easy to read format.
func (ml *MoveList) String() string {
	builder := strings.Builder{}
	for _, move := range ml.moves {
		builder.WriteString(move.String())
		builder.WriteString("\n")
	}

	builder.WriteString("\n")

	return builder.String()
}

// func (ml *MoveList) sort() {
// 	cmp := func(i, j Move) bool {
// 		if i.GetPieceType() == j.GetPieceType() {
// 			if i.From == j.getFrom() {
// 				return i.GetTo() < j.GetTo()
// 			}
// 			return i.GetFrom() < j.GetFrom()
// 		}
// 		return i.GetPieceType() < j.GetPieceType()
// 	}
// 	quickSort(ml.moves, cmp)
// }
