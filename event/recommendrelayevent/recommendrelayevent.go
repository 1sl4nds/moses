package recommendrelayevent

import (
	"github.com/1sl4nds/moses/event"
)

const Kind = 2

func New() *event.Event {
	return event.New(Kind, "")
}
