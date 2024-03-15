package channelhidemessageevent

import (
	"github.com/1sl4nds/moses/event"
)

// Kind for hiding messages in a channel
const Kind = 43

// New TODO
func New() *event.Event {
	return event.New(Kind, "")
}
