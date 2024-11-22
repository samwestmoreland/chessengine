package position

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	"github.com/samwestmoreland/chessengine/internal/piece"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/internal/utils"
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

	// Map pieces to their indices and colour bitboards
	pieceMap := map[rune]struct {
		index      piece.Piece
		colourBits piece.Piece
	}{
		'P': {piece.Wp, piece.Wa},
		'N': {piece.Wn, piece.Wa},
		'B': {piece.Wb, piece.Wa},
		'R': {piece.Wr, piece.Wa},
		'Q': {piece.Wq, piece.Wa},
		'K': {piece.Wk, piece.Wa},
		'p': {piece.Bp, piece.Ba},
		'n': {piece.Bn, piece.Ba},
		'b': {piece.Bb, piece.Ba},
		'r': {piece.Br, piece.Ba},
		'q': {piece.Bq, piece.Ba},
		'k': {piece.Bk, piece.Ba},
	}

	var square sq.Square

	for _, char := range posStr {
		// Handle pieces
		if info, isPiece := pieceMap[char]; isPiece {
			occ[info.index] = bb.SetBit(occ[info.index], square)
			occ[info.colourBits] = bb.SetBit(occ[info.colourBits], square)
			square++

			continue
		}

		// Handle empty squares
		if char >= '1' && char <= '8' {
			square += sq.Square(char - '0')

			continue
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
				log.Printf("checking square %s for piece %d", sq.Stringify(square), i)
				if bb.GetBit(occ, square) {
					occupied = true

					utils.WriteOrDie(" "+piece.Piece(byte(i)).String(), output)

					break
				}
			}

			if !occupied {
				utils.WriteOrDie(" .", output)
			}
		}

		utils.WriteOrDie("\n", output)
	}

	utils.WriteOrDie(fmt.Sprintf("\nside to move: %s\n", sideToString(p.WhiteToMove)), output)
	utils.WriteOrDie(fmt.Sprintf("castling rights: %s\n", castlingRightsToString(p.CastlingRights)), output)
	utils.WriteOrDie(fmt.Sprintf("en passant square: %s\n", sq.Stringify(p.EnPassantSquare)), output)
}

func (p *Position) Copy() *Position {
	newOccupancy := make([]bb.Bitboard, len(p.Occupancy))
	copy(newOccupancy, p.Occupancy)

	return &Position{
		Occupancy:       newOccupancy,
		WhiteToMove:     p.WhiteToMove,
		CastlingRights:  p.CastlingRights,
		EnPassantSquare: p.EnPassantSquare,
		HalfMoveClock:   p.HalfMoveClock,
		FullMoveNumber:  p.FullMoveNumber,
	}
}

func (p *Position) MakeMove(source, target sq.Square, movePiece piece.Piece) *Position {
	ret := p.Copy()

	ret.ClearSquare(target)
	ret.ClearSquare(source)

	ret.PlacePiece(target, movePiece)

	return ret
}

func (p *Position) ClearSquare(square sq.Square) {
	for i := piece.Wp; i < piece.Bk; i++ {
		if bb.GetBit(p.Occupancy[i], square) {
			p.Occupancy[i] = bb.ClearBit(p.Occupancy[i], square)

			colour, err := i.Colour()
			if err != nil {
				panic(err)
			}

			if colour == piece.Black {
				p.Occupancy[piece.Ba] = bb.ClearBit(p.Occupancy[piece.Ba], square)
			} else {
				p.Occupancy[piece.Wa] = bb.ClearBit(p.Occupancy[piece.Wa], square)
			}

			break
		}
	}
}

func (p *Position) PlacePiece(square sq.Square, pieceToPlace piece.Piece) {
	p.Occupancy[pieceToPlace] = bb.SetBit(p.Occupancy[pieceToPlace], square)

	colour, err := pieceToPlace.Colour()
	if err != nil {
		panic(err)
	}

	if colour == piece.Black {
		p.Occupancy[piece.Ba] = bb.SetBit(p.Occupancy[piece.Ba], square)
	} else {
		p.Occupancy[piece.Wa] = bb.SetBit(p.Occupancy[piece.Wa], square)
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
