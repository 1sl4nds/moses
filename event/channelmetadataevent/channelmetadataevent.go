package channelmetadataevent

import "github.com/1sl4nds/moses/event"

// Kind for setting channel metadata
const Kind = 41

// New creates a new channel metadata event.
func New() *event.Event {
	return event.New(Kind, "")
}
