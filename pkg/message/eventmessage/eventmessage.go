package eventmessage

import (
	"fmt"

	"github.com/1sl4nds/moses/message"
	"github.com/1sl4nds/moses/pkg/event"
)

const Type = "EVENT"

// New creates a new EventMessage.
func New(subscriptionID string, evt *event.Event) message.Message {
	if len(subscriptionID) > 64 {
		panic(fmt.Errorf("invalid subscription id"))
	}
	return message.New(Type, subscriptionID, evt)
}
