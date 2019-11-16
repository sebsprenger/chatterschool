package client

import "github.com/sebsprenger/chatterschool/shared"

type OutputFormatter interface {
	FormatMessage(msg shared.Message)
}
