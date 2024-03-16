package eventtag

import "github.com/1sl4nds/moses/pkg/tag"

const Type = "a"

func New(naddr string) tag.Tag {
	return tag.New(Type, naddr)
}
