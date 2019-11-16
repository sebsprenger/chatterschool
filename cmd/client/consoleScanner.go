package main

import (
	"bufio"
	"fmt"
	"os"
)

var errQuit = fmt.Errorf("User wants to quit")

type ConsoleChatScanner struct {
	scanner *bufio.Scanner
}

func NewConsoleChatScanner() ConsoleChatScanner {
	return ConsoleChatScanner{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (reader ConsoleChatScanner) readFromConsole() (string, error) {
	for {
		reader.scanner.Scan()
		text := reader.scanner.Text()

		switch text {
		case "":
			continue
		case "/quit":
			return "", errQuit
		}

		return text, nil
	}
}
