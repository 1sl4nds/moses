package createorupdateproductevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

// Kind for creating or updating a product
const Kind = 30018

// New creates a new CreateOrUpdateProductEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
