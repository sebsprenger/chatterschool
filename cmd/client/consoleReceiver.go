package main

import (
	"fmt"

	"github.com/sebsprenger/chatterschool/shared"
)

type ConsoleReceiver struct {
}

func (formatter ConsoleReceiver) FormatMessage(msg shared.Message) {
	fmt.Printf("%s: %s\n", msg.Sender, msg.Text)
}
