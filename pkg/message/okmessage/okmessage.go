package okmessage

import "github.com/1sl4nds/moses/pkg/message"

const Type = "OK"

// New creates a new OkMessage.
func New(eventID string, ok bool, status string) message.Message {
	return message.New(Type, eventID, ok, status)
}
