package contactsevent

import (
	"github.com/1sl4nds/moses/event"
)

// Kind for contact List and petnames
const Kind = 3

// New creates a new contacts event.
func New() *event.Event {
	return event.New(Kind, "")
}
