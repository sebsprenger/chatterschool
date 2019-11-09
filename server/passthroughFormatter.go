package server

import (
	"github.com/sebsprenger/chatterschool/shared"
)

type PassThroughFormatter interface {
	Modify(msg shared.Message) shared.Message
}
