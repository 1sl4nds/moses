package closemessage

import "github.com/1sl4nds/moses/pkg/message"

// Type TBD
const Type = "CLOSE"

// New creates a new CloseMessage.
func New() message.Message {
	return message.New(Type)
}
