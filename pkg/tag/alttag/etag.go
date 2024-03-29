package eventidtag

import "github.com/1sl4nds/moses/pkg/tag"

const Type = "e"

type Options struct {
	Marker string
}

func New(eventID string, relayURL string, opt *Options) tag.Tag {
	t := tag.New(Type, eventID, relayURL)
	if opt == nil {
		return t
	}
	if opt.Marker != "" {
		t.Push(opt.Marker)
	}
	return t
}
