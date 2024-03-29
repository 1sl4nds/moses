package channelcreationevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

// Kind for creating new channels
const Kind = 40

// New creates a new channel creation event.
func New() *event.Event {
	return event.New(Kind, "")
}
