package utils

import "io"

func WriteOrDie(s string, writer io.Writer) {
	_, err := writer.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}
