package main

import (
	"bufio"
	"bytes"
	"os"
)

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

		resp, quit := handleCommand(cmd)

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

func handleCommand(cmd string) (*bytes.Buffer, bool) {
	var resp bytes.Buffer

	var quit bool

	// Process command
	switch cmd {
	case "uci\n":
		_, err := resp.WriteString("id name Toto Chess Engine\n")
		if err != nil {
			panic(err)
		}

		_, err = resp.WriteString("id author Sam Westmoreland\n")
		if err != nil {
			panic(err)
		}
	case "quit\n":
		_, err := resp.WriteString("Bye!\n")
		if err != nil {
			panic(err)
		}

		quit = true
	default:
		_, err := resp.WriteString("Unknown command\n")
		if err != nil {
			panic(err)
		}
	}

	return &resp, quit
}
