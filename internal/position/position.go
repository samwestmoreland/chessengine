package position

import (
	"errors"
	"fmt"
	"strings"

	"github.com/samwestmoreland/chessengine/internal/board"
	"github.com/samwestmoreland/chessengine/internal/moves"
	"github.com/samwestmoreland/chessengine/internal/piece"
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

func NewStartingPosition() *Position {
	fen := NewFEN()

	return getPositionFromFEN(fen)
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

func (p *Position) GetTurn() board.Colour {
	return p.Turn
}

func (p *Position) GetWhitePieces() map[board.Square]Piece {
	return p.White
}

func (p *Position) GetBlackPieces() map[board.Square]Piece {
	return p.Black
}

func (p *Position) GetAllMovesConcurrent(turn board.Colour) (moves.MoveList, error) {
	var numPieces int

	var pieces map[board.Square]Piece

	if turn == board.White {
		pieces = p.White
		numPieces = len(p.White)
		log.Debugf("Getting moves for %d white pieces\n", len(p.White))
	} else if turn == board.Black {
		pieces = p.Black
		numPieces = len(p.Black)
		log.Debugf("Getting moves for %d black pieces\n", len(p.Black))
	}

	ret := moves.MoveList{}
	movesChan := make(chan moves.MoveList, numPieces)

	// start goroutines to get moves for each piece
	for _, pce := range pieces {
		go func(myPiece Piece, myChan chan<- moves.MoveList, pos Position) {
			pieceMoves, err := myPiece.GetMoves(&pos)
			if err != nil {
				log.Errorf("Failed to get moves for white piece %v: %v\n", myPiece.Type(), err)
			}

			myChan <- pieceMoves
		}(pce, movesChan, *p)
	}

	for range pieces {
		ret.AddMoveList(<-movesChan)
	}

	return ret, nil
}

// GetAllMovesSerial returns all possible moves for the current position
// without any concurrency. This is just for benchmarking, really.
func (p *Position) GetAllMovesSerial(turn board.Colour) (moves.MoveList, error) {
	var movs moves.MoveList

	var pieces *map[board.Square]Piece

	if turn == board.White {
		pieces = &p.White
	} else if turn == board.Black {
		pieces = &p.Black
	}

	for _, piece := range *pieces {
		pieceMoves, err := piece.GetMoves(p)
		if err != nil {
			return movs, fmt.Errorf("failed to get moves for black piece %v: %w", piece.Type(), err)
		}

		movs.AddMoveList(pieceMoves)
	}

	return movs, nil
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
func (p *Position) GetAllPossibleMoves() (moves.MoveList, error) {
	ret := moves.MoveList{}
	if p.White == nil || p.Black == nil {
		return ret, fmt.Errorf("position is not valid: %w", ErrInvalidPosition)
	}

	if p.Turn == board.White {
		return p.getWhiteMoves()
	}

	if p.Turn == board.Black {
		return p.getBlackMoves()
	}

	return ret, fmt.Errorf("position is not valid: %w", ErrInvalidPosition)
}

func (p *Position) getWhiteMoves() (moves.MoveList, error) {
	var movs moves.MoveList

	for square, piece := range p.White {
		pieceMoves, err := piece.GetMoves(p)
		if err != nil {
			return movs,
				fmt.Errorf("failed to get moves for white piece %v on square %s: %w",
					piece.Type(), square.String(), err)
		}

		movs.AddMoveList(pieceMoves)
	}

	return movs, nil
}

func (p *Position) getBlackMoves() (moves.MoveList, error) {
	ret := moves.MoveList{}

	for square, piece := range p.Black {
		pieceMoves, err := piece.GetMoves(p)
		if err != nil {
			return ret,
				fmt.Errorf("failed to get moves for black piece %v on square %s: %w",
					piece.Type(), square.String(), err)
		}

		ret.AddMoveList(pieceMoves)
	}

	return ret, nil
}

// getPiece returns the piece at the given square, or an error if the square is invalid.
func (p *Position) getPiece(square board.Square) (Piece, error) {
	if !square.Valid() {
		return nil, fmt.Errorf("invalid square: %s: %w", square.String(), board.ErrInvalidSquare)
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

// squareIsOccupied returns true if the given square is occupied by a piece, and the colour of the piece on that square.
func (p *Position) squareIsOccupied(square board.Square) (bool, board.Colour) {
	if _, ok := p.White[square]; ok {
		return true, board.White
	}

	if _, ok := p.Black[square]; ok {
		return true, board.Black
	}

	return false, board.Unknown
}

func (p *Position) GetNumPiecesForColour(pieceType piece.Type, col board.Colour) int {
	var pieces *map[board.Square]Piece

	if col == board.White {
		pieces = &p.White
	} else if col == board.Black {
		pieces = &p.Black
	}

	var ret int

	for _, piece := range *pieces {
		if piece.Type() == pieceType {
			ret++
		}
	}

	return ret
}

func (p *Position) GetPiecesForColour(pieceType piece.Type, col board.Colour) map[board.Square]Piece {
	var pieces *map[board.Square]Piece

	if col == board.White {
		pieces = &p.White
	} else if col == board.Black {
		pieces = &p.Black
	}

	ret := make(map[board.Square]Piece)

	for square, piece := range *pieces {
		if piece.Type() == pieceType {
			ret[square] = piece
		}
	}

	return ret
}
