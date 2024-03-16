package identifiertag

import "github.com/1sl4nds/moses/tag"

const Type = "d"

func New(identifier string) tag.Tag {
	return tag.New(Type, identifier)
}
