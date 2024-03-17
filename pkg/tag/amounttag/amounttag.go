package amounttag

import (
	"github.com/1sl4nds/moses/pkg/tag"
)

const Type = "amount"

func New(amount int) *tag.Tag {
	return &tag.Tag{
		Type,
		amount,
	}
}
