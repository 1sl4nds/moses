package noticemessage

import "github.com/1sl4nds/moses/pkg/message"

const Type = "NOTICE"

// New creates a new NoticeMessage.
func New() message.Message {
	return message.New(Type)
}
