package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/samwestmoreland/chessengine/magic_data"
	"github.com/samwestmoreland/chessengine/src/bitboard"
	sq "github.com/samwestmoreland/chessengine/src/squares"
	"github.com/samwestmoreland/chessengine/src/tables"
)

var (
	rookRelevantBits = [64]int{
		12, 11, 11, 11, 11, 11, 11, 12,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		12, 11, 11, 11, 11, 11, 11, 12,
	}

	bishopRelevantBits = [64]int{
		6, 5, 5, 5, 5, 5, 5, 6,
		5, 5, 5, 5, 5, 5, 5, 5,
		5, 5, 7, 7, 7, 7, 5, 5,
		5, 5, 7, 9, 9, 7, 5, 5,
		5, 5, 7, 9, 9, 7, 5, 5,
		5, 5, 7, 7, 7, 7, 5, 5,
		5, 5, 5, 5, 5, 5, 5, 5,
		6, 5, 5, 5, 5, 5, 5, 6,
	}
)

type MagicEntry struct {
	Square string `json:"square"`
	Magic  string `json:"magic"` // as hex string
	Shift  int    `json:"shift"`
	Mask   string `json:"mask"` // as hex string
}

type MagicData struct {
	Rook struct {
		Magics []MagicEntry `json:"magics"`
	} `json:"rook"`
	Bishop struct {
		Magics []MagicEntry `json:"magics"`
	} `json:"bishop"`
	Metadata struct {
		Generated string `json:"generated"`
		Version   string `json:"version"`
	} `json:"metadata"`
}

const (
	bishop = iota
	rook
)

var versionString = strings.TrimSpace(magic_data.VersionString)

func main() {
	rookMagics := generateMagics(rook)
	bishopMagics := generateMagics(bishop)

	today := time.Now().Format("2006-01-02 15:04:05")

	magicData := &MagicData{
		Rook: struct {
			Magics []MagicEntry `json:"magics"`
		}{
			Magics: rookMagics,
		},
		Bishop: struct {
			Magics []MagicEntry `json:"magics"`
		}{
			Magics: bishopMagics,
		},
		Metadata: struct {
			Generated string `json:"generated"`
			Version   string `json:"version"`
		}{
			Generated: today,
			Version:   versionString,
		},
	}

	data, err := json.MarshalIndent(magicData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("magic_data/magics.json", data, 0644); err != nil {
		log.Fatal(err)
	}
}

func generateMagics(piece int) []MagicEntry {
	if !(piece == rook || piece == bishop) {
		log.Fatal("piece must be rook or bishop")
	}

	magics := []MagicEntry{}

	for square := 0; square < 64; square++ {
		bestTableSize := math.MaxInt64
		var bestMagic uint64
		var bestShift int

		var relevantBits int
		if piece == rook {
			relevantBits = rookRelevantBits[square]
		} else {
			relevantBits = bishopRelevantBits[square]
		}

		for shift := 64 - relevantBits; shift < 64-relevantBits+4; shift++ {
			for attempt := 0; attempt < 1000000; attempt++ {
				magicCandidate := bitboard.GenerateSparseRandomUint64()

				if works, tableSize := testMagicCandidate(magicCandidate, square, shift, piece, relevantBits); works {
					if tableSize < bestTableSize {
						bestTableSize = tableSize
						bestMagic = magicCandidate
						bestShift = shift
					}
				}
			}

			if bestMagic != 0 {
				break
			}
		}

		entry := MagicEntry{
			Square: sq.Stringify(square),
			Magic:  fmt.Sprintf("%016x", bestMagic),
			Shift:  bestShift,
		}

		var mask uint64
		if piece == rook {
			mask = tables.MaskRookAttacks(square)
		} else {
			mask = tables.MaskBishopAttacks(square)
		}

		entry.Mask = fmt.Sprintf("%016x", mask)

		magics = append(magics, entry)
	}

	return magics
}

func testMagicCandidate(magicCandidate uint64, square, shift, piece, relevantBits int) (bool, int) {
	var maxIndex int

	var numBlockerConfigs = 1 << relevantBits

	used := make(map[int]uint64) // index -> possible moves

	for blockerConfigIndex := 0; blockerConfigIndex < numBlockerConfigs; blockerConfigIndex++ {
		var attacks uint64
		if piece == rook {
			attacks = tables.MaskRookAttacks(square)
		} else {
			attacks = tables.MaskBishopAttacks(square)
		}

		blockerConfig := bitboard.SetOccupancy(blockerConfigIndex, attacks)

		var actualMoves uint64
		if piece == rook {
			actualMoves = tables.RookAttacksOnTheFly(square, blockerConfig)
		} else {
			actualMoves = tables.BishopAttacksOnTheFly(square, blockerConfig)
		}

		hashResult := (blockerConfig * magicCandidate) >> shift
		index := int(hashResult)

		if index > maxIndex {
			maxIndex = index
		}

		if _, ok := used[index]; !ok { // new index
			used[index] = actualMoves
		} else {
			// check if the actual moves are the same
			if used[index] != actualMoves {
				return false, 0
			}
		}
	}

	return true, maxIndex + 1
}
