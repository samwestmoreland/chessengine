package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
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
		_, err := resp.WriteString("id name Toto Chess Engine\n")
		if err != nil {
			panic(err)
		}

		_, err = resp.WriteString("id author Sam Westmoreland\n")
		if err != nil {
			panic(err)
		}
	case "quit":
		_, err := resp.WriteString("Bye!\n")
		if err != nil {
			panic(err)
		}

		quit = true
	case "position":
		posResp, posQuit := handlePositionCmd(cmd)
		quit = posQuit

		_, err := resp.Write(posResp.Bytes())
		if err != nil {
			panic(err)
		}
	default:
		_, err := resp.WriteString("Unknown command\n")
		if err != nil {
			panic(err)
		}
	}

	return &resp, quit
}

func handlePositionCmd(cmd *command) (*bytes.Buffer, bool) {
	var resp bytes.Buffer

	var quit bool

	switch cmd.args[0] {
	case "startpos":
		_, err := resp.WriteString("set up starting position\n")
		if err != nil {
			panic(err)
		}
	default:
		_, err := resp.WriteString("Unknown command\n")
		if err != nil {
			panic(err)
		}
	}

	return &resp, quit
}
