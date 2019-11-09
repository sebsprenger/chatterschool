package main

import (
	"fmt"
	"strings"

	"github.com/sebsprenger/chatterschool/shared"
)

type MyPassThroughFormatter struct {
}

func (formatter MyPassThroughFormatter) Modify(msg shared.Message) shared.Message {
	fmt.Printf("eavesdropping: %s\n", msg)

	// censoring
	msg.Text = strings.ReplaceAll(msg.Text, "scheisse", "s******e")
	msg.Text = strings.ReplaceAll(msg.Text, "scheiße", "s*****e")

	return msg
}
