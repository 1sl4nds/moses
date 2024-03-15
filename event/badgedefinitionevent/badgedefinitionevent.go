package badgedefinitionevent

import (
	"github.com/1sl4nds/moses/event"
	"github.com/1sl4nds/moses/tag"
)

const Kind = 30009

func New(content string, tags ...tag.Tag) *event.Event {
	return event.New(Kind, content, tags...)
}
