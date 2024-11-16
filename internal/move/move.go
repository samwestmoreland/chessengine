package move

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

func Encode(source, target, piece, promotionPiece, capture, doublePush, enPassant, castling uint32) Move {
	return Move(
		(source) |
			(target << 6) |
			(piece << 12) |
			(promotionPiece << 16) |
			(capture << 20) |
			(doublePush << 21) |
			(enPassant << 22) |
			(castling << 23))
}

func (m Move) Source() int {
	return int(m & 0b111111)
}

func (m Move) Target() int {
	return int((m >> 6) & 0b111111)
}

// MoveBuilder type for debugging and testing
type MoveBuilder struct {
	source       int
	target       int
	piece        int
	promotion    int
	isCapture    bool
	isDoublePush bool
	isEnPassant  bool
	isCastling   bool
}

func NewMove() *MoveBuilder {
	return &MoveBuilder{}
}

func (b *MoveBuilder) From(square int) *MoveBuilder {
	b.source = square
	return b
}

func (b *MoveBuilder) To(square int) *MoveBuilder {
	b.target = square
	return b
}

func (b *MoveBuilder) Piece(p int) *MoveBuilder {
	b.piece = p
	return b
}

func (b *MoveBuilder) Promotion(p int) *MoveBuilder {
	b.promotion = p
	return b
}

func (b *MoveBuilder) Capture() *MoveBuilder {
	b.isCapture = true
	return b
}

func (b *MoveBuilder) DoublePush() *MoveBuilder {
	b.isDoublePush = true
	return b
}

func (b *MoveBuilder) EnPassant() *MoveBuilder {
	b.isEnPassant = true
	return b
}

func (b *MoveBuilder) Castling() *MoveBuilder {
	b.isCastling = true
	return b
}

func (b *MoveBuilder) Build() Move {
	var move uint32

	// Pack all the fields into the move
	move |= uint32(b.source)
	move |= uint32(b.target) << 6
	move |= uint32(b.piece) << 12
	move |= uint32(b.promotion) << 16

	if b.isCapture {
		move |= captureMask
	}
	if b.isDoublePush {
		move |= doublePushMask
	}
	if b.isEnPassant {
		move |= enPassantMask
	}
	if b.isCastling {
		move |= castlingMask
	}

	return Move(move)
}
