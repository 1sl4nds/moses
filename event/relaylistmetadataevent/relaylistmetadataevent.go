package relaylistmetadataevent

import "github.com/1sl4nds/moses/event"

// Kind for managing relay list metadata
const Kind = 10002

// New creates a new RelayListMetadataEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
