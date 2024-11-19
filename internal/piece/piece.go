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

var pieceToString = map[Piece]string{
	Wp: "P",
	Wn: "N",
	Wb: "B",
	Wr: "R",
	Wq: "Q",
	Wk: "K",
	Bp: "p",
	Bn: "n",
	Bb: "b",
	Br: "r",
	Bq: "q",
	Bk: "k",
}

func (p Piece) String() string {
	if s, ok := pieceToString[p]; ok {
		return s
	}

	return ""
}

type Colour uint8

const (
	White Colour = iota
	Black
)
