package badgenametag

import "github.com/1sl4nds/moses/tag"

const Type = "name"

func New(name string) tag.Tag {
	return tag.Tag{
		Type,
		name,
	}
}
