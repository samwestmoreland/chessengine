package piece

// Type represents the type of a piece.
type Type int

const (
	KingType   Type = iota // KingType = 0
	QueenType              // QueenType = 1
	RookType               // RookType = 2
	BishopType             // BishopType = 3
	KnightType             // KnightType = 4
	PawnType               // PawnType = 5
	NoneType               // NoneType = 6
)

func (t Type) String() string {
	switch t {
	case KingType:
		return "King"
	case QueenType:
		return "Queen"
	case RookType:
		return "Rook"
	case BishopType:
		return "Bishop"
	case KnightType:
		return "Knight"
	case PawnType:
		return "Pawn"
	case NoneType:
		return "None"
	default:
		return "Unknown"
	}
}
