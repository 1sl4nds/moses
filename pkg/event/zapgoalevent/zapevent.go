package zapevent

import "github.com/1sl4nds/moses/pkg/event"

// Kind for performing a Zap action
const Kind = 9735

// New creates a new ZapEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
