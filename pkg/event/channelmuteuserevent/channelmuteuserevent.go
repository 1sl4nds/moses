package channelmuteuserevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

// Kind for muting a user in a channel
const Kind = 44

// New creates a new ChannelMuteUserEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
