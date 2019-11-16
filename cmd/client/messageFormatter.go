package main

import (
	"github.com/sebsprenger/chatterschool/shared"
)

type MessageFormatter struct {
	sender string
}

func NewMessageFormatter(senderName string) MessageFormatter {
	return MessageFormatter{
		sender: senderName,
	}
}

func (formatter MessageFormatter) CreateMessage(input string) shared.Message {
	input = formatter.changeInput(input)
	return formatter.buildMessage(input)
}

func (formatter MessageFormatter) changeInput(input string) string {
	if input == "/shrug" {
		input = `¯\_(ツ)_/¯`
	}
	return input
}

func (formatter MessageFormatter) buildMessage(input string) shared.Message {
	msg := shared.Message{
		Text:   input,
		Sender: formatter.sender,
	}
	return msg
}
