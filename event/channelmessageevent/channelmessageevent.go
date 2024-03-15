package channelmessagevent

import "github.com/1sl4nds/moses/event"

// Event for posting messages in a channel
const Kind = 42

// New creates a new channel message event.
func New() *event.Event {
	return event.New(Kind, "")
}
