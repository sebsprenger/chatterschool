package main

import (
	"github.com/sebsprenger/chatterschool/shared"
)

type MyInputFormatter struct {
	sender string
}

func (formatter MyInputFormatter) CreateMessage(input string) shared.Message {
	if input == "/shrug" {
		input = `¯\_(ツ)_/¯`
	}

	msg := shared.Message{
		Text:   input,
		Sender: formatter.sender,
	}
	return msg
}
