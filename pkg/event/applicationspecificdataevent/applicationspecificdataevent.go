package applicationspecificdataevent

import (
	"github.com/1sl4nds/moses/pkg/event"
)

const Kind = 30078

func New(content string) *event.Event {
	return event.New(Kind, content)
}
