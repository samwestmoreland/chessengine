package tables

import (
	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
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

func isAFile(square int) bool {
	return (1 << square & notAFile) == 0
}

func isHFile(square int) bool {
	return (1 << square & notHFile) == 0
}

func isABFile(square int) bool {
	return (1 << square & notABFile) == 0
}

func isGHFile(square int) bool {
	return (1 << square & notGHFile) == 0
}

func ConstGenerator() {
	var board bb.Bitboard

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			// Convert rank and file into a square
			square := rank*8 + file

			if file > 1 {
				board = bb.SetBit(board, square)
			}
		}
	}

	bb.PrintBoard(board)
}
