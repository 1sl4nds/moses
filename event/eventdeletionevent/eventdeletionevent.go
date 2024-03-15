package eventdeletionevent

import (
	"github.com/1sl4nds/moses/event"
)

// Kind for deleting events
const Kind = 5

// New creates a new event deletion event.
func New(content string) *event.Event {
	return event.New(Kind, content)
}
