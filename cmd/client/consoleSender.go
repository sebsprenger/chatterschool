package main

import (
	"fmt"

	"github.com/sebsprenger/chatterschool/client"
)

type ConsoleSender struct {
	scanner          ConsoleChatScanner
	messageFormatter MessageFormatter
}

func (consoleChat ConsoleSender) SendChatMessagesTo(client *client.ChatClient) {
	for {
		input, err := consoleChat.scanner.readFromConsole()
		if err != nil {
			break
		}

		msg := consoleChat.messageFormatter.CreateMessage(input)

		err = client.Send(msg)
		if err != nil {
			fmt.Printf("Error while sending message: %s\n", err)
		}
	}
}
