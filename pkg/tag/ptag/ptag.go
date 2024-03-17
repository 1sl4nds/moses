package ptag

import "github.com/1sl4nds/moses/pkg/tag"

const Type = "p"

func New(pubKeyStr string, relayURL string, petname string) tag.Tag {
	return tag.New(Type, pubKeyStr, relayURL, petname)
}
