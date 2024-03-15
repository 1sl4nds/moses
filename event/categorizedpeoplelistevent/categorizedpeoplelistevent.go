package categorizedpeoplelistevent

import (
	"github.com/1sl4nds/moses/event"
	"github.com/1sl4nds/moses/tag"
)

// Kind for managing categorized people list
const Kind = 30000

// New creates a new CategorizedPeopleListEvent.
func New(content string, tags ...tag.Tag) *event.Event {
	return event.New(Kind, content, tags...)
}
