package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
	"time"

	"github.com/samwestmoreland/chessengine/src/engine"
	"github.com/samwestmoreland/chessengine/src/position"
	"github.com/samwestmoreland/chessengine/src/tables"
)

type command struct {
	name string
	args []string
}

var lookupTable tables.Lookup

func main() {
	if err := initialiseLookupTables(&lookupTable); err != nil {
		log.Fatal(err)
	}

	runEngine()
}

func runEngine() {
	eng := engine.NewEngine()

	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	for {
		_, err := writer.WriteString("engine ready\n")
		if err != nil {
			panic(err)
		}

		writer.Flush()

		cmd, err := reader.ReadString('\n')
		if err != nil {
			_, err := writer.WriteString("Error reading input\n")
			if err != nil {
				panic(err)
			}

			writer.Flush()

			continue
		}

		parsed := parseCmd(cmd)
		resp, quit := handleCommand(parsed, eng)

		_, err = resp.WriteTo(writer)
		if err != nil {
			panic(err)
		}

		writer.Flush()

		if quit {
			break
		}
	}
}

func parseCmd(cmd string) *command {
	cmd = strings.Trim(cmd, "\n")

	parts := strings.Split(cmd, " ")

	return &command{
		name: parts[0],
		args: parts[1:],
	}
}

func handleCommand(cmd *command, eng *engine.Engine) (*bytes.Buffer, bool) {
	var resp bytes.Buffer

	var quit bool

	// Process command
	switch cmd.name {
	case "uci":
		mustWrite(&resp, "id name Toto Chess Engine\n")
		mustWrite(&resp, "id author Sam Westmoreland\n")
	case "quit", "exit", "bye", "q":
		mustWrite(&resp, "bye!\n")

		quit = true
	case "position":
		posResp := handlePositionCmd(cmd, eng)

		_, err := resp.Write(posResp.Bytes())
		if err != nil {
			panic(err)
		}
	case "isready":
		mustWrite(&resp, "readyok\n")
	default:
		mustWrite(&resp, "unknown command\n")
	}

	return &resp, quit
}

func handlePositionCmd(cmd *command, eng *engine.Engine) *bytes.Buffer {
	var resp bytes.Buffer

	if len(cmd.args) == 0 || cmd.args == nil {
		mustWrite(&resp, "too few arguments. expected `position <fen>` or `position startpos`\n")

		return &resp
	}

	if cmd.args[0] == "startpos" {
		mustWrite(&resp, "set up starting position\n")

		return &resp
	}

	// Try to parse FEN
	fen, err := position.ParseFEN(strings.Join(cmd.args, " "))
	if err != nil {
		mustWrite(&resp, "invalid FEN\n")

		return &resp
	}

	pos := position.NewPositionFromFEN(fen)
	eng.SetPosition(pos)

	return &resp
}

func mustWrite(buf *bytes.Buffer, s string) {
	_, err := buf.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func initialiseLookupTables(lookup *tables.Lookup) error {
	log.Print("initialising lookup tables")
	start := time.Now()

	if err := tables.InitialiseLookupTables(lookup); err != nil {
		return err
	}

	log.Printf("initialised lookup tables in %s", time.Since(start))

	return nil
}
