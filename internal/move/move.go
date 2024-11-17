package move

import (
	"github.com/samwestmoreland/chessengine/internal/piece"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

// Moves are represented as bitmasks:
//
//	0000 0000 0000 0000 0011 1111    source square
//	0000 0000 0000 1111 1100 0000    target square
//	0000 0000 1111 0000 0000 0000    piece
//	0000 1111 0000 0000 0000 0000    promotion piece
//	0001 0000 0000 0000 0000 0000    capture flag
//	0010 0000 0000 0000 0000 0000    double push flag
//	0100 0000 0000 0000 0000 0000    en passant flag
//	1000 0000 0000 0000 0000 0000    castling flag
type Move uint32

func Encode(
	source, target sq.Square,
	movePiece, promotionPiece piece.Piece,
	capture, doublePush, enPassant, castling uint32,
) Move {
	return Move(
		uint32(source) |
			(uint32(target) << 6) |
			(uint32(movePiece) << 12) |
			(uint32(promotionPiece) << 16) |
			(capture << 20) |
			(doublePush << 21) |
			(enPassant << 22) |
			(castling << 23))
}

func (m Move) String() string {
	ret := sq.Stringify(m.Source()) + sq.Stringify(m.Target())

	if m.PromotionPiece() != piece.NoPiece {
		ret += m.PromotionPiece().String()
	}

	return ret
}

func (m Move) Source() sq.Square {
	val := m & 0x3f

	return sq.Square(byte(val))
}

func (m Move) Target() sq.Square {
	val := (m >> 6) & 0x3f

	return sq.Square(byte(val))
}

func (m Move) Piece() piece.Piece {
	val := (m >> 12) & 0xf

	return piece.Piece(byte(val)) // bytehis an alias for uint8
}

func (m Move) PromotionPiece() piece.Piece {
	val := (m >> 16) & 0xf

	return piece.Piece(byte(val))
}

func (m Move) IsCapture() bool {
	return (m >> 20) == 1
}

func (m Move) IsDoublePush() bool {
	return (m >> 21) == 1
}

func (m Move) IsEnPassant() bool {
	return (m >> 22) == 1
}

func (m Move) IsCastling() bool {
	return (m >> 23) == 1
}

// Builder type for debugging and testing.
type Builder struct {
	source       sq.Square
	target       sq.Square
	piece        piece.Piece
	promotion    piece.Piece
	isCapture    bool
	isDoublePush bool
	isEnPassant  bool
	isCastling   bool
}

func NewMove() *Builder {
	return &Builder{
		source:       sq.NoSquare,
		target:       sq.NoSquare,
		piece:        piece.NoPiece,
		promotion:    piece.NoPiece,
		isCapture:    false,
		isDoublePush: false,
		isEnPassant:  false,
		isCastling:   false,
	}
}

func (b *Builder) From(square sq.Square) *Builder {
	b.source = square

	return b
}

func (b *Builder) To(square sq.Square) *Builder {
	b.target = square

	return b
}

func (b *Builder) Piece(p piece.Piece) *Builder {
	b.piece = p

	return b
}

func (b *Builder) Promotion(p piece.Piece) *Builder {
	b.promotion = p

	return b
}

func (b *Builder) Capture() *Builder {
	b.isCapture = true

	return b
}

func (b *Builder) DoublePush() *Builder {
	b.isDoublePush = true

	return b
}

func (b *Builder) EnPassant() *Builder {
	b.isEnPassant = true

	return b
}

func (b *Builder) Castling() *Builder {
	b.isCastling = true

	return b
}

func (b *Builder) Build() Move {
	var move Move

	// Pack all the fields into the move
	move |= Move(uint32(b.source))
	move |= Move(uint32(b.target) << 6)
	move |= Move(uint32(b.piece) << 12)
	move |= Move(uint32(b.promotion) << 16)

	if b.isCapture {
		move |= (1 << 20)
	}

	if b.isDoublePush {
		move |= (1 << 21)
	}

	if b.isEnPassant {
		move |= (1 << 22)
	}

	if b.isCastling {
		move |= (1 << 23)
	}

	return move
}
