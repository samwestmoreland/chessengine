package main

import (
	"os"

	"github.com/integrii/flaggy"
	"github.com/samwestmoreland/chessengine/src/position"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var version = "v0.0.0"

func main() {
	flaggy.SetName("Chess Engine")
	flaggy.SetDescription("A chess engine written in Go")
	flaggy.SetVersion(version)

	var fenFlag string
	flaggy.String(&fenFlag, "f", "fen", "A FEN string to parse")

	flaggy.Parse()
	log.Infof("Starting chess engine")
	_, err := position.ParseFEN(fenFlag)
	if err != nil {
		log.Errorf("Error parsing FEN: %s", err)
		os.Exit(1)
	}
}
