package position

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/internal/piece"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

type Position struct {
	Occupancy       []bb.Bitboard // Pieces of both colours
	WhiteToMove     bool
	CastlingRights  uint8
	EnPassantSquare sq.Square
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
	parts := strings.Split(fen, " ")
	// parts[0]: position string
	// parts[1]: turn to move
	// parts[2]: castling rights
	// parts[3]: en passant square
	// parts[4]: halfmove clock
	// parts[5]: fullmove number

	if len(parts) != 6 {
		return nil, fmt.Errorf("FEN must have 6 parts, got %d", len(parts))
	}

	occ, err := parsePositionString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse position string: %w", err)
	}

	whiteToMove, err := parseSideToMove(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse side to move: %w", err)
	}

	castlingRights, err := parseCastlingRights(parts[2])
	if err != nil {
		return nil, fmt.Errorf("failed to parse castling rights: %w", err)
	}

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
		WhiteToMove:     whiteToMove,
		CastlingRights:  castlingRights,
		EnPassantSquare: enpassant,
		HalfMoveClock:   byte(halfMoveClock),
		FullMoveNumber:  byte(fullMoveNumber),
	}, nil
}

func parseSideToMove(side string) (bool, error) {
	if side == "w" {
		return true, nil
	} else if side == "b" {
		return false, nil
	}

	return false, fmt.Errorf("invalid side to move: %s", side)
}

func (p *Position) IsOccupied(square sq.Square) bool {
	return bb.GetBit(p.Occupancy[piece.Wa], square) || bb.GetBit(p.Occupancy[piece.Ba], square)
}

func parseEnPassantSquare(square string) (sq.Square, error) {
	if square == "-" {
		return sq.Square(sq.NoSquare), nil
	}

	squareInt, err := sq.ParseString(square)
	if err != nil {
		return 0, fmt.Errorf("failed to parse en passant square: %w", err)
	}

	return squareInt, nil
}

func parsePositionString(posStr string) ([]bb.Bitboard, error) {
	occ := make([]bb.Bitboard, 15)

	var square sq.Square

	for i := range len(posStr) {
		switch posStr[i] {
		case 'P':
			occ[piece.Wp] = bb.SetBit(occ[piece.Wp], square)
			occ[piece.Wa] = bb.SetBit(occ[piece.Wa], square)
			square++
		case 'N':
			occ[piece.Wn] = bb.SetBit(occ[piece.Wn], square)
			occ[piece.Wa] = bb.SetBit(occ[piece.Wa], square)
			square++
		case 'B':
			occ[piece.Wb] = bb.SetBit(occ[piece.Wb], square)
			occ[piece.Wa] = bb.SetBit(occ[piece.Wa], square)
			square++
		case 'R':
			occ[piece.Wr] = bb.SetBit(occ[piece.Wr], square)
			occ[piece.Wa] = bb.SetBit(occ[piece.Wa], square)
			square++
		case 'Q':
			occ[piece.Wq] = bb.SetBit(occ[piece.Wq], square)
			occ[piece.Wa] = bb.SetBit(occ[piece.Wa], square)
			square++
		case 'K':
			occ[piece.Wk] = bb.SetBit(occ[piece.Wk], square)
			occ[piece.Wa] = bb.SetBit(occ[piece.Wa], square)
			square++
		case 'p':
			occ[piece.Bp] = bb.SetBit(occ[piece.Bp], square)
			occ[piece.Ba] = bb.SetBit(occ[piece.Ba], square)
			square++
		case 'n':
			occ[piece.Bn] = bb.SetBit(occ[piece.Bn], square)
			occ[piece.Ba] = bb.SetBit(occ[piece.Ba], square)
			square++
		case 'b':
			occ[piece.Bb] = bb.SetBit(occ[piece.Bb], square)
			occ[piece.Ba] = bb.SetBit(occ[piece.Ba], square)
			square++
		case 'r':
			occ[piece.Br] = bb.SetBit(occ[piece.Br], square)
			occ[piece.Ba] = bb.SetBit(occ[piece.Ba], square)
			square++
		case 'q':
			occ[piece.Bq] = bb.SetBit(occ[piece.Bq], square)
			occ[piece.Ba] = bb.SetBit(occ[piece.Ba], square)
			square++
		case 'k':
			occ[piece.Bk] = bb.SetBit(occ[piece.Bk], square)
			occ[piece.Ba] = bb.SetBit(occ[piece.Ba], square)
			square++
		case '1':
			square++
		case '2':
			square += 2
		case '3':
			square += 3
		case '4':
			square += 4
		case '5':
			square += 5
		case '6':
			square += 6
		case '7':
			square += 7
		case '8':
			square += 8
		}
	}

	if square != 64 {
		return nil, fmt.Errorf("expected 64 squares, got %d", square)
	}

	return occ, nil
}

// parseCastlingRights parses a string of castling rights into a uint8
// 1000: K
// 0100: Q
// 0010: k
// 0001: q
func parseCastlingRights(castlingRights string) (uint8, error) {
	expectedLength := 4

	if len(castlingRights) > expectedLength {
		return 0, fmt.Errorf("expected castling rights to be %d characters, got %d", expectedLength, len(castlingRights))
	}

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

	return ret, nil
}

func (p *Position) Print(output io.Writer) {
	for rank := range 8 {
		for file := range 8 {
			square := sq.Square(byte(rank*8 + file))
			occupied := false

			for i, occ := range p.Occupancy {
				if bb.GetBit(occ, square) {
					occupied = true

					if _, err := output.Write([]byte(" " + piece.Piece(byte(i)).String())); err != nil {
						panic(err)
					}

					break
				}
			}

			if !occupied {
				if _, err := output.Write([]byte(" .")); err != nil {
					panic(err)
				}
			}
		}

		if _, err := output.Write([]byte("\n")); err != nil {
			panic(err)
		}
	}

	if _, err := output.Write([]byte(fmt.Sprintf(
		"\nside to move: %s\n", sideToString(p.WhiteToMove),
	))); err != nil {
		panic(err)
	}

	if _, err := output.Write([]byte(fmt.Sprintf(
		"castling rights: %s\n", castlingRightsToString(p.CastlingRights),
	))); err != nil {
		panic(err)
	}

	if _, err := output.Write([]byte(fmt.Sprintf(
		"en passant square: %s\n", sq.Stringify(p.EnPassantSquare),
	))); err != nil {
		panic(err)
	}
}

func sideToString(whiteToMove bool) string {
	if whiteToMove {
		return "white"
	}

	return "black"
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
