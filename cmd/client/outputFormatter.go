package main

import (
	"fmt"

	"github.com/sebsprenger/chatterschool/shared"
)

type MyOutputFormatter struct {
}

func (formatter MyOutputFormatter) FormatMessage(msg shared.Message) string {
	return fmt.Sprintf("%s: %s", msg.Sender, msg.Text)
}
