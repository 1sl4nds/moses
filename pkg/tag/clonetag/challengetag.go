package challengetag

import "github.com/1sl4nds/moses/tag"

const Type = "challenge"

func New(challenge string) tag.Tag {
	return tag.Tag{
		Type,
		challenge,
	}
}
