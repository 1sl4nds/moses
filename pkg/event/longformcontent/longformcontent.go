package longformcontent

import "github.com/1sl4nds/moses/pkg/event"

// Kind for posting long-form content
const Kind = 30023

// New creates a new LongFormContentEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
