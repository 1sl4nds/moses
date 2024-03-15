package profilebadges

import (
	"github.com/1sl4nds/moses/event"
)

// Kind for profile badges management
const Kind = 30008

// New creates a new ProfileBadgesEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
