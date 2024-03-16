package clientauthenticationevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

// Event for client authentication process
const Kind = 22242

// New creates a new ClientAuthenticationEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
