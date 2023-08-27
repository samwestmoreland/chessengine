package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/samwestmoreland/chessengine/src/position"
)

type command struct {
	name string
	args []string
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	for {
		_, err := writer.WriteString("engine ready\n") // send initial ready message
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
		resp, quit := handleCommand(parsed)

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

func handleCommand(cmd *command) (*bytes.Buffer, bool) {
	var resp bytes.Buffer

	var quit bool

	// Process command
	switch cmd.name {
	case "uci":
		mustWrite(&resp, "id name Toto Chess Engine\n")
		mustWrite(&resp, "id author Sam Westmoreland\n")
	case "quit":
		mustWrite(&resp, "bye!\n")

		quit = true
	case "position":
		posResp := handlePositionCmd(cmd)

		_, err := resp.Write(posResp.Bytes())
		if err != nil {
			panic(err)
		}
	default:
		mustWrite(&resp, "unknown command\n")
	}

	return &resp, quit
}

func handlePositionCmd(cmd *command) *bytes.Buffer {
	var resp bytes.Buffer

	if len(cmd.args) == 0 || cmd.args == nil {
		mustWrite(&resp, "too few arguments. expected `position <fen>` or `position startpos`\n")

		return &resp
	}

	if cmd.args[0] == "startpos" {
		mustWrite(&resp, "set up starting position\n")

		return &resp
	}

	// try to parse FEN
	_, err := position.ParseFEN(cmd.args[0])
	if err != nil {
		mustWrite(&resp, "invalid FEN\n")
	}

	return &resp
}

func mustWrite(buf *bytes.Buffer, s string) {
	_, err := buf.WriteString(s)
	if err != nil {
		panic(err)
	}
}
