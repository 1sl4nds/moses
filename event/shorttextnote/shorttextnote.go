package shorttextnote

import (
	"github.com/1sl4nds/moses/event"
)

const Kind = 1

func New(content string) *event.Event {
	return event.New(Kind, content)
}
