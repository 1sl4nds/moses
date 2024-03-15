package badgedesctag

import "github.com/1sl4nds/moses/tag"

const Type = "description"

// New tag representing the image, the meaning behind the badge, or the reason of it's issuance.
func New(description string) tag.Tag {
	return tag.Tag{
		Type,
		description,
	}
}

type Options struct {
	Description string
}
