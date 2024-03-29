package connectevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

// Kind for Nostr client connection
const Kind = 24133

// New creates a new NostrConnectEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
