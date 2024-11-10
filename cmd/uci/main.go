package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/samwestmoreland/chessengine/internal/engine"
	"github.com/samwestmoreland/chessengine/internal/movegen"
	"github.com/samwestmoreland/chessengine/internal/position"
)

type UCI struct {
	engine   *engine.Engine
	position *position.Position
	writer   *bufio.Writer
	reader   *bufio.Reader
}

func NewUCI(writer *bufio.Writer, reader *bufio.Reader) (*UCI, error) {
	if err := movegen.Initialise(); err != nil {
		return nil, err
	}

	eng, err := engine.NewEngine()
	if err != nil {
		return nil, err
	}

	return &UCI{
		engine: eng,
		writer: writer,
		reader: reader,
	}, nil
}

func (u *UCI) Run() error {
	for {
		if _, err := u.writer.WriteString("engine ready\n"); err != nil {
			return err
		}
		u.writer.Flush()

		cmdStr, err := u.reader.ReadString('\n')
		if err != nil {
			if _, err := u.writer.WriteString("Error reading input\n"); err != nil {
				return err
			}
			u.writer.Flush()
			continue
		}

		cmd := parseCmd(cmdStr)
		resp, quit := u.handleCommand(cmd)

		if _, err := resp.WriteTo(u.writer); err != nil {
			return err
		}
		u.writer.Flush()

		if quit {
			break
		}
	}
	return nil
}

func (u *UCI) handleCommand(cmd *command) (*bytes.Buffer, bool) {
	var resp bytes.Buffer
	var quit bool

	switch cmd.name {
	case "uci":
		resp.WriteString("id name Toto Chess Engine\n")
		resp.WriteString("id author Sam Westmoreland\n")
	case "quit", "exit", "bye", "q":
		resp.WriteString("bye!\n")

		quit = true
	case "position":
		u.handlePositionCmd(cmd, &resp)

	case "isready":
		resp.WriteString("readyok\n")
	case "ponder":
		bestMove := u.engine.Search(u.position)
		resp.WriteString(bestMove + "\n")
	default:
		resp.WriteString("unknown command\n")
	}

	return &resp, quit
}

func (u *UCI) handlePositionCmd(cmd *command, resp *bytes.Buffer) {
	if len(cmd.args) == 0 {
		resp.WriteString("too few arguments. expected `position <fen>` or `position startpos`\n")
		return
	}

	if cmd.args[0] == "startpos" {
		pos, err := position.NewPosition()
		if err != nil {
			panic(err)
		}

		u.position = pos
		resp.WriteString("set up starting position\n")

		pos.Print(os.Stdout)

		return
	}

	// Handle FEN string case...
}

func main() {
	uci, err := NewUCI(
		bufio.NewWriter(os.Stdout),
		bufio.NewReader(os.Stdin),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := uci.Run(); err != nil {
		log.Fatal(err)
	}
}

type command struct {
	name string
	args []string
}

func parseCmd(cmd string) *command {
	cmd = strings.Trim(cmd, "\n")

	parts := strings.Split(cmd, " ")

	return &command{
		name: parts[0],
		args: parts[1:],
	}
}
