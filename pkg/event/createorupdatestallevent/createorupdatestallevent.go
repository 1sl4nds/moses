package createorupdatestallevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

// Kind for creating or updating a stall
const Kind = 30017

// New creates a new CreateOrUpdateStallEvent.
func New() *event.Event {
	return &event.Event{}
}
