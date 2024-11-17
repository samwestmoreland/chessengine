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

	bb "github.com/samwestmoreland/chessengine/internal/bitboard"
	sq "github.com/samwestmoreland/chessengine/internal/squares"
	"github.com/samwestmoreland/chessengine/internal/tables"
	"github.com/samwestmoreland/chessengine/magic"
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

	if err := os.WriteFile("magic/magics.json", data, 0600); err != nil {
		log.Fatal(err)
	}
}

func generateMagics(piece int) ([]magic.Entry, uint64) {
	if piece == rook {
		log.Println("Generating rook magics")
	} else if piece == bishop {
		log.Println("Generating bishop magics")
	} else {
		log.Fatal("Piece must be rook or bishop")
	}

	numWorkers := runtime.GOMAXPROCS(0)
	log.Printf("Using %d workers", numWorkers)

	squares := make(chan sq.Square, 64)

	var wg sync.WaitGroup

	var totalTableSize atomic.Uint64

	magics := make([]magic.Entry, 64)

	for w := 0; w < numWorkers; w++ {
		log.Printf("Spawning worker %d", w)

		wg.Add(1)

		go func() {
			defer wg.Done()

			for square := range squares {
				var bestTableSize uint64 = math.MaxUint64

				var bestMagic uint64

				var bestShift int

				var relevantBits int
				if piece == rook {
					relevantBits = rookRelevantBits[square]
				} else {
					relevantBits = bishopRelevantBits[square]
				}

				for shift := 64 - relevantBits; shift < 64-relevantBits+4; shift++ {
					for attempt := 0; attempt < 10000000; attempt++ {
						magicCandidate := bb.GenerateSparseRandomUint64()

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

				totalTableSize.Add(bestTableSize)

				entry := magic.Entry{
					Square: sq.Stringify(square),
					Magic:  fmt.Sprintf("%016x", bestMagic),
					Shift:  bestShift,
					Mask:   "",
				}

				var mask bb.Bitboard
				if piece == rook {
					mask = tables.MaskRookAttacks(square)
				} else {
					mask = tables.MaskBishopAttacks(square)
				}

				entry.Mask = fmt.Sprintf("%016x", mask)

				magics[square] = entry

				log.Printf("Found magic for square %s", sq.Stringify(square))
			}
		}()
	}

	// Feed squares to workers
	for square := uint8(0); square < 64; square++ {
		squares <- sq.Square(square)
	}

	close(squares)

	// Wait for all workers to finish
	wg.Wait()

	return magics, totalTableSize.Load()
}

func testMagicCandidate(magicCandidate uint64, square sq.Square, shift, piece, relevantBits int) (bool, uint64) {
	var maxIndex uint64

	var numBlockerConfigs = 1 << relevantBits

	used := make(map[uint64]bb.Bitboard) // index -> possible moves

	for blockerConfigIndex := 0; blockerConfigIndex < numBlockerConfigs; blockerConfigIndex++ {
		var attacks bb.Bitboard
		if piece == rook {
			attacks = tables.MaskRookAttacks(square)
		} else {
			attacks = tables.MaskBishopAttacks(square)
		}

		blockerConfig := bb.SetOccupancy(blockerConfigIndex, attacks)

		var actualMoves bb.Bitboard
		if piece == rook {
			actualMoves = tables.RookAttacksOnTheFly(square, blockerConfig)
		} else {
			actualMoves = tables.BishopAttacksOnTheFly(square, blockerConfig)
		}

		hashResult := (uint64(blockerConfig) * magicCandidate) >> shift
		index := hashResult

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

func formatTableSize(numEntries uint64) string {
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
