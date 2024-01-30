package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetReader() *bufio.Reader {
	return reader
}

func GetReaderPrompt(prompt string) string {
	fmt.Printf("%s: ", prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error: %s \n", err)
		GetReaderPrompt(prompt)
	}

	input = strings.TrimSpace(input)

	return input
}
