package position

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
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

type Position struct {
	Occupancy       []bb.Bitboard // Pieces of both colours
	WhiteToMove     bool
	CastlingRights  uint8
	EnPassantSquare uint8
	HalfMoveClock   uint8
	FullMoveNumber  uint8
}

func NewPosition() (*Position, error) {
	pos, err := NewPositionFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		return nil, fmt.Errorf("failed to create new position: %w", err)
	}

	return pos, nil
}

func NewPositionFromFEN(fen string) (*Position, error) {
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

	occ, err := parsePositionString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse position string: %w", err)
	}

	castlingRights := parseCastlingRights(parts[2])

	enpassant, err := parseEnPassantSquare(parts[3])
	if err != nil {
		return nil, fmt.Errorf("failed to parse en passant square: %w", err)
	}

	halfMoveClock, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse half move clock: %w", err)
	}

	fullMoveNumber, err := strconv.Atoi(parts[5])
	if err != nil {
		return nil, fmt.Errorf("failed to parse full move number: %w", err)
	}

	return &Position{
		Occupancy:       occ,
		WhiteToMove:     parts[1] == "w",
		CastlingRights:  castlingRights,
		EnPassantSquare: uint8(enpassant),
		HalfMoveClock:   uint8(halfMoveClock),
		FullMoveNumber:  uint8(fullMoveNumber),
	}, nil
}

func (p *Position) IsOccupied(square int) bool {
	return bb.GetBit(p.Occupancy[A], square) || bb.GetBit(p.Occupancy[a], square)
}

func parseEnPassantSquare(square string) (uint8, error) {
	if square == "-" {
		return uint8(sq.NoSquare), nil
	}

	squareInt, err := sq.ToInt(square)
	if err != nil {
		return 0, fmt.Errorf("failed to parse en passant square: %w", err)
	}

	return uint8(squareInt), nil
}

func parsePositionString(posStr string) ([]bb.Bitboard, error) {
	occ := make([]bb.Bitboard, 14)

	sq := 0

	for i := 0; i < len(posStr); i++ {
		switch posStr[i] {
		case 'P':
			occ[P] = bb.SetBit(occ[P], sq)
			occ[A] = bb.SetBit(occ[A], sq)
			sq++
		case 'N':
			occ[N] = bb.SetBit(occ[N], sq)
			occ[A] = bb.SetBit(occ[A], sq)
			sq++
		case 'B':
			occ[B] = bb.SetBit(occ[B], sq)
			occ[A] = bb.SetBit(occ[A], sq)
			sq++
		case 'R':
			occ[R] = bb.SetBit(occ[R], sq)
			occ[A] = bb.SetBit(occ[A], sq)
			sq++
		case 'Q':
			occ[Q] = bb.SetBit(occ[Q], sq)
			occ[A] = bb.SetBit(occ[A], sq)
			sq++
		case 'K':
			occ[K] = bb.SetBit(occ[K], sq)
			occ[A] = bb.SetBit(occ[A], sq)
			sq++
		case 'p':
			occ[p] = bb.SetBit(occ[p], sq)
			occ[a] = bb.SetBit(occ[a], sq)
			sq++
		case 'n':
			occ[n] = bb.SetBit(occ[n], sq)
			occ[a] = bb.SetBit(occ[a], sq)
			sq++
		case 'b':
			occ[b] = bb.SetBit(occ[b], sq)
			occ[a] = bb.SetBit(occ[a], sq)
			sq++
		case 'r':
			occ[r] = bb.SetBit(occ[r], sq)
			occ[a] = bb.SetBit(occ[a], sq)
			sq++
		case 'q':
			occ[q] = bb.SetBit(occ[q], sq)
			occ[a] = bb.SetBit(occ[a], sq)
			sq++
		case 'k':
			occ[k] = bb.SetBit(occ[k], sq)
			occ[a] = bb.SetBit(occ[a], sq)
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
		return nil, fmt.Errorf("expected 64 squares, got %d", sq)
	}

	return occ, nil
}

// parseCastlingRights parses a string of castling rights into a uint8
// 1000: K
// 0100: Q
// 0010: k
// 0001: q
func parseCastlingRights(castlingRights string) uint8 {

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

func (s *Position) Print(output io.Writer) {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			occupied := false
			for i, occ := range s.Occupancy {
				if bb.GetBit(occ, square) {
					occupied = true
					output.Write([]byte(fmt.Sprintf(" %s", GetPieceType(i))))
					break
				}
			}

			if !occupied {
				output.Write([]byte(fmt.Sprintf(" .")))
			}
		}

		output.Write([]byte(fmt.Sprintf("\n")))
	}

	output.Write([]byte(fmt.Sprintf("\nside to move: %s\n", sideToString(s.WhiteToMove))))
	output.Write([]byte(fmt.Sprintf("castling rights: %s\n", castlingRightsToString(s.CastlingRights))))
	output.Write([]byte(fmt.Sprintf("en passant square: %s\n", sq.Stringify(int(s.EnPassantSquare)))))
}

func sideToString(whiteToMove bool) string {
	if whiteToMove {
		return "white"
	} else {
		return "black"
	}
}

func castlingRightsToString(castlingRights uint8) string {
	sb := strings.Builder{}

	if castlingRights&8 == 8 {
		sb.WriteString("K")
	}

	if castlingRights&4 == 4 {
		sb.WriteString("Q")
	}

	if castlingRights&2 == 2 {
		sb.WriteString("k")
	}

	if castlingRights&1 == 1 {
		sb.WriteString("q")
	}

	return sb.String()
}
