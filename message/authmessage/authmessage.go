package authmessage

import (
	"github.com/1sl4nds/moses/event"
	"github.com/1sl4nds/moses/message"
)

// Type for message "AUTH"
const Type = "AUTH"

// New creates a new "AUTH" message
func New(challenge string, evt *event.Event) message.Message {
	if evt != nil {
		return message.New(Type, evt)
	}
	return message.New(Type, challenge)
}
