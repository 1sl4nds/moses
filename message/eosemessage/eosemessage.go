package eosemessage

import (
	"github.com/1sl4nds/moses/message"
)

const Type = "EOSE"

func New() message.Message {
	return message.New(Type)
}
