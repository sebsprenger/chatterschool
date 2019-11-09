package client

import "github.com/sebsprenger/chatterschool/shared"

type InputFormatter interface {
	CreateMessage(input string) shared.Message
}
