package eventrepostsevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

const Kind = 6

func New() *event.Event {
	return event.New(Kind, "")
}
