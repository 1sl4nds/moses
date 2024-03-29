package badgedefinitionevent

import (
	"github.com/1sl4nds/moses/pkg/event"
	"github.com/1sl4nds/moses/pkg/tag"
)

const Kind = 30009

func New(content string, tags ...tag.Tag) *event.Event {
	return event.New(Kind, content, tags...)
}
