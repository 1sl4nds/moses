package badgeawardevent

import (
	"github.com/1sl4nds/moses/event"
)

// Kind for awarding badges to users
const Kind = 8

// New creates a new badge award event.
func New(content string) *event.Event {
	return event.New(Kind, content)
}
