package position

import (
	"errors"
	"fmt"
	"strings"

	"github.com/samwestmoreland/chessengine/src/board"
	"github.com/samwestmoreland/chessengine/src/moves"
	"github.com/sirupsen/logrus"
)

var (
	log                = logrus.New()
	ErrInvalidPosition = errors.New("invalid position")
)

// A Position represents a chess position.
type Position struct {
	White map[board.Square]Piece
	Black map[board.Square]Piece
	Turn  board.Colour
}

// NewPosition returns a new Position.
func NewPositionFromFEN(fen *FEN) *Position {
	position := getPositionFromFEN(fen)

	return position
}

func (p *Position) String() string {
	// Print an ascii representation of the board.
	var ret string

	for rank := 8; rank >= 1; rank-- {
		for file := 1; file <= 8; file++ {
			square := board.Square{File: file, Rank: rank}

			piece, err := p.getPiece(square)
			if err != nil {
				log.Errorf("Failed to get piece on square %s: %v\n", square.String(), err)

				return ""
			}

			if piece == nil {
				ret += ". "

				continue
			}

			if piece.GetColour() == board.White {
				ret += piece.Type().Letter() + " "
			} else if piece.GetColour() == board.Black {
				ret += strings.ToLower(piece.Type().Letter()) + " "
			}
		}

		ret += "\n"
	}

	return ret
}

func NewPosition(turn board.Colour, pieces []Piece) *Position {
	whitePieces := make(map[board.Square]Piece)
	blackPieces := make(map[board.Square]Piece)
	ret := &Position{Turn: turn, White: whitePieces, Black: blackPieces}

	for _, piece := range pieces {
		if !piece.GetCurrentSquare().Valid() {
			log.Errorf("Failed to add piece %v to square %v\n", piece.Type(), piece.GetCurrentSquare())

			continue
		}

		if piece.GetColour() == board.White {
			whitePieces[piece.GetCurrentSquare()] = piece
		} else if piece.GetColour() == board.Black {
			blackPieces[piece.GetCurrentSquare()] = piece
		}
	}

	ret.White = whitePieces
	ret.Black = blackPieces

	return ret
}

func getPositionFromFEN(fen *FEN) *Position {
	log.Debugf("Creating position from FEN: %s\n", fen.String())
	white, black := getPiecePositionsFromFEN(fen)
	ret := Position{White: white, Black: black, Turn: fen.Colour}

	return &ret
}

func getPiecePositionsFromFEN(fen *FEN) (map[board.Square]Piece, map[board.Square]Piece) {
	white := make(map[board.Square]Piece)
	black := make(map[board.Square]Piece)

	for _, square := range board.Squares {
		piece, err := fen.GetPiece(square)
		if err != nil {
			continue
		}

		if piece == nil {
			continue
		}

		square := square

		if piece.GetColour() == board.White {
			white[square] = piece
		} else {
			black[square] = piece
		}
	}

	return white, black
}

// GetAllPossibleMoves returns all possible moves for the current position.
func (p *Position) GetAllPossibleMoves() ([]moves.Move, error) {
	var moves []moves.Move
	if p.White == nil || p.Black == nil {
		return moves, fmt.Errorf("position is not valid: %w", ErrInvalidPosition)
	}

	if p.Turn == board.White {
		return p.getWhiteMoves()
	}

	if p.Turn == board.Black {
		return p.getBlackMoves()
	}

	return moves, fmt.Errorf("position is not valid: %w", ErrInvalidPosition)
}

func (p *Position) getWhiteMoves() ([]moves.Move, error) {
	var moves []moves.Move

	for square, piece := range p.White {
		pieceMoves, err := piece.GetMoves(p)
		if err != nil {
			return moves,
				fmt.Errorf("failed to get moves for white piece %v on square %s: %w",
					piece.Type(), square.String(), err)
		}

		moves = append(moves, pieceMoves...)
	}

	return moves, nil
}

func (p *Position) getBlackMoves() ([]moves.Move, error) {
	var moves []moves.Move

	for square, piece := range p.Black {
		pieceMoves, err := piece.GetMoves(p)
		if err != nil {
			return moves,
				fmt.Errorf("failed to get moves for black piece %v on square %s: %w",
					piece.Type(), square.String(), err)
		}

		moves = append(moves, pieceMoves...)
	}

	return moves, nil
}

// getPiece returns the piece at the given square, or an error if the square is invalid.
func (p *Position) getPiece(square board.Square) (Piece, error) {
	if !square.Valid() {
		return nil, fmt.Errorf("invalid square: %s", square.String())
	}

	if piece, ok := p.White[square]; ok {
		return piece, nil
	}

	if piece, ok := p.Black[square]; ok {
		return piece, nil
	}

	var piece Piece

	return piece, nil
}
