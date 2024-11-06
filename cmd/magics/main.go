package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/samwestmoreland/chessengine/magic"
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

const (
	bishop = iota
	rook
)

var versionString = strings.TrimSpace(magic.VersionString)

func main() {
	// bar := progressbar.NewOptions(128, progressbar.OptionSetTheme(progressbar.Theme{
	// 	Saucer:        "=",
	// 	SaucerHead:    ">",
	// 	SaucerPadding: " ",
	// 	BarStart:      "[",
	// 	BarEnd:        "]",
	// }))

	rookMagics, rookTableSize := generateMagics(rook)
	bishopMagics, bishopTableSize := generateMagics(bishop)

	today := time.Now().Format("2006-01-02 15:04:05")

	magicData := magic.Data{
		Rook: magic.RookData{
			Magics:         rookMagics,
			TotalTableSize: formatTableSize(rookTableSize),
		},
		Bishop: magic.BishopData{
			Magics:         bishopMagics,
			TotalTableSize: formatTableSize(bishopTableSize),
		},
		Metadata: magic.Metadata{
			TotalTableSize: formatTableSize(rookTableSize + bishopTableSize),
			Generated:      today,
			Version:        versionString,
		},
	}

	data, err := json.MarshalIndent(magicData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("magic/magics.json", data, 0644); err != nil {
		log.Fatal(err)
	}
}

func generateMagics(piece int) ([]magic.Entry, int) {
	if piece == rook {
		log.Println("Generating rook magics")
	} else if piece == bishop {
		log.Println("Generating bishop magics")
	} else {
		log.Fatal("Piece must be rook or bishop")
	}

	numWorkers := runtime.GOMAXPROCS(0)
	log.Printf("Using %d workers", numWorkers)
	squares := make(chan int, 64)
	var wg sync.WaitGroup

	var totalTableSize atomic.Int64

	magics := make([]magic.Entry, 64)

	for w := 0; w < numWorkers; w++ {
		log.Printf("Spawning worker %d", w)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for square := range squares {
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

				if bestMagic == 0 {
					log.Printf("Failed to find a magic for square %d", square)
				}

				totalTableSize.Add(int64(bestTableSize))

				entry := magic.Entry{
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

				magics[square] = entry
			}
		}()
	}

	// Feed squares to workers
	for square := 0; square < 64; square++ {
		squares <- square
	}
	close(squares)

	// Wait for all workers to finish
	wg.Wait()

	return magics, int(totalTableSize.Load())
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

func formatTableSize(numEntries int) string {
	bytes := numEntries * 8 // 8 bytes per uint64

	switch {
	case bytes < 1024:
		return fmt.Sprintf("%d B", bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2f KB", float64(bytes)/1024)
	default:
		return fmt.Sprintf("%.2f MB", float64(bytes)/(1024*1024))
	}
}
