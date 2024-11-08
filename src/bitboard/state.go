package bitboard

import (
	"fmt"
	"log"
	"strings"
)

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

func GetPieceType(pieceInt int) string {
	switch pieceInt {
	case P:
		return "P"
	case N:
		return "N"
	case B:
		return "B"
	case R:
		return "R"
	case Q:
		return "Q"
	case K:
		return "K"
	case p:
		return "p"
	case n:
		return "n"
	case b:
		return "b"
	case r:
		return "r"
	case q:
		return "q"
	case k:
		return "k"
	default:
		return ""
	}
}

type State struct {
	Occupancy       []uint64 // Pieces of both colours
	WhiteToMove     bool
	CastlingRights  uint8
	EnPassantSquare uint8
	HalfMoveClock   uint8
	FullMoveNumber  uint8
}

func NewState() *State {
	return &State{
		Occupancy:       make([]uint64, 14),
		WhiteToMove:     false,
		CastlingRights:  0,
		EnPassantSquare: 0,
		HalfMoveClock:   0,
		FullMoveNumber:  0,
	}
}

func NewStateFromFEN(fen string) (*State, error) {
	// e.g.	"r3k1nr/pp2pp1p/6p1/1b1q4/8/2b2N2/PP2QPPP/R1B2RK1 w kq - 0 13"
	parts := strings.Split(fen, " ")
	// parts[0]: position string
	// parts[1]: turn to move
	// parts[2]: castling rights
	// parts[3]: en passant square
	// parts[4]: halfmove clock
	// parts[5]: fullmove number

	if len(parts) != 6 {
		return nil, fmt.Errorf("FEN must have 6 parts")
	}

	occ := parsePositionString(parts[0])
	castlingRights := parseCastlingRights(parts[2])

	return &State{
		Occupancy:       occ,
		WhiteToMove:     parts[1] == "w",
		CastlingRights:  castlingRights,
		EnPassantSquare: 0,
	}, nil
}

func parsePositionString(posStr string) []uint64 {
	occ := make([]uint64, 14)

	sq := 0

	for i := 0; i < len(posStr); i++ {
		switch posStr[i] {
		case 'P':
			occ[P] |= 1 << sq
			occ[A] |= 1 << sq
			sq++
		case 'N':
			occ[N] |= 1 << sq
			occ[A] |= 1 << sq
			sq++
		case 'B':
			occ[B] |= 1 << sq
			occ[A] |= 1 << sq
			sq++
		case 'R':
			occ[R] |= 1 << sq
			occ[A] |= 1 << sq
			sq++
		case 'Q':
			occ[Q] |= 1 << sq
			occ[A] |= 1 << sq
			sq++
		case 'K':
			occ[K] |= 1 << sq
			occ[A] |= 1 << sq
			sq++
		case 'p':
			occ[p] |= 1 << sq
			occ[a] |= 1 << sq
			sq++
		case 'n':
			occ[n] |= 1 << sq
			occ[a] |= 1 << sq
			sq++
		case 'b':
			occ[b] |= 1 << sq
			occ[a] |= 1 << sq
			sq++
		case 'r':
			occ[r] |= 1 << sq
			occ[a] |= 1 << sq
			sq++
		case 'q':
			occ[q] |= 1 << sq
			occ[a] |= 1 << sq
			sq++
		case 'k':
			occ[k] |= 1 << sq
			occ[a] |= 1 << sq
			sq++
		case '1':
			sq += 1
		case '2':
			sq += 2
		case '3':
			sq += 3
		case '4':
			sq += 4
		case '5':
			sq += 5
		case '6':
			sq += 6
		case '7':
			sq += 7
		case '8':
			sq += 8
		}
	}

	if sq != 64 {
		log.Fatalf("Expected 64 squares in FEN, got %d", sq)
	}

	return occ
}

func parseCastlingRights(castlingRights string) uint8 {

	// 1000: K
	// 0100: Q
	// 0010: k
	// 0001: q

	var ret uint8

	if strings.Contains(castlingRights, "K") {
		ret |= 8
	}

	if strings.Contains(castlingRights, "Q") {
		ret |= 4
	}

	if strings.Contains(castlingRights, "k") {
		ret |= 2
	}

	if strings.Contains(castlingRights, "q") {
		ret |= 1
	}

	return ret
}

func (s *State) Print() {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			occupied := false
			for i, occ := range s.Occupancy {
				if GetBit(occ, square) {
					occupied = true
					fmt.Printf(" %s", GetPieceType(i))
					break
				}
			}

			if !occupied {
				fmt.Printf(" .")
			}
		}

		fmt.Printf("\n")
	}
}
