package mutelistevent

import (
	"github.com/1sl4nds/moses/event"
)

// Kind for managing a mute list
const Kind = 10000

// New creates a new MuteListEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
