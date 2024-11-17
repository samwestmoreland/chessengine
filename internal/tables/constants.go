package tables

import (
	"os"

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
)

// notAFile is const represeting the board:
//
//    8  0 1 1 1 1 1 1 1
//    7  0 1 1 1 1 1 1 1
//    6  0 1 1 1 1 1 1 1
//    5  0 1 1 1 1 1 1 1
//    4  0 1 1 1 1 1 1 1
//    3  0 1 1 1 1 1 1 1
//    2  0 1 1 1 1 1 1 1
//    1  0 1 1 1 1 1 1 1
//       a b c d e f g h
//

const notAFile bb.Bitboard = 18374403900871474942

// notHFile is const represeting the board:
//
//    8  1 1 1 1 1 1 1 0
//    7  1 1 1 1 1 1 1 0
//    6  1 1 1 1 1 1 1 0
//    5  1 1 1 1 1 1 1 0
//    4  1 1 1 1 1 1 1 0
//    3  1 1 1 1 1 1 1 0
//    2  1 1 1 1 1 1 1 0
//    1  1 1 1 1 1 1 1 0
//       a b c d e f g h
//

const notHFile bb.Bitboard = 9187201950435737471

// notABFile is const represeting the board:
//
//	8  0 0 1 1 1 1 1 1
//	7  0 0 1 1 1 1 1 1
//	6  0 0 1 1 1 1 1 1
//	5  0 0 1 1 1 1 1 1
//	4  0 0 1 1 1 1 1 1
//	3  0 0 1 1 1 1 1 1
//	2  0 0 1 1 1 1 1 1
//	1  0 0 1 1 1 1 1 1
//	   a b c d e f g h
//

const notABFile bb.Bitboard = 18229723555195321596

// notGHFile is const represeting the board:
//
//	8  1 1 1 1 1 1 0 0
//	7  1 1 1 1 1 1 0 0
//	6  1 1 1 1 1 1 0 0
//	5  1 1 1 1 1 1 0 0
//	4  1 1 1 1 1 1 0 0
//	3  1 1 1 1 1 1 0 0
//	2  1 1 1 1 1 1 0 0
//	1  1 1 1 1 1 1 0 0
//	   a b c d e f g h
//

const notGHFile bb.Bitboard = 4557430888798830399

func isAFile(square sq.Square) bool {
	return (1 << square & notAFile) == 0
}

func isHFile(square sq.Square) bool {
	return (1 << square & notHFile) == 0
}

func isABFile(square sq.Square) bool {
	return (1 << square & notABFile) == 0
}

func isGHFile(square sq.Square) bool {
	return (1 << square & notGHFile) == 0
}

func ConstGenerator() {
	var board bb.Bitboard

	for rank := range 8 {
		for file := range 8 {
			// Convert rank and file into a square
			square := sq.Square(byte(rank*8 + file))

			if file > 1 {
				board = bb.SetBit(board, square)
			}
		}
	}

	bb.PrintBoard(board, os.Stdout)
}
