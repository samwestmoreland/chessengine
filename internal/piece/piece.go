package piece

type Piece uint8

const (
	NoPiece Piece = iota
	Wp
	Wn
	Wb
	Wr
	Wq
	Wk
	Bp
	Bn
	Bb
	Br
	Bq
	Bk
	Wa // All white
	Ba // All black
)

func (p Piece) String() string {
	switch p {
	case Wp:
		return "P"
	case Wn:
		return "N"
	case Wb:
		return "B"
	case Wr:
		return "R"
	case Wq:
		return "Q"
	case Wk:
		return "K"
	case Bp:
		return "p"
	case Bn:
		return "n"
	case Bb:
		return "b"
	case Br:
		return "r"
	case Bq:
		return "q"
	case Bk:
		return "k"
	default:
		return ""
	}
}
