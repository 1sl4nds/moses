package requestmessage

import (
	"github.com/1sl4nds/moses/message"
)

const Type = "REQ"

// New creates a new RequestMessage with the provided subscription ID and filters.
func New(sid string, filter ...*Filter) message.Message {
	msg := message.New(Type, sid)
	for _, f := range filter {
		msg.Push(f)
	}
	return msg
}
