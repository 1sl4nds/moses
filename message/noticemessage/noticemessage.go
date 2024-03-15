package noticemessage

import "github.com/1sl4nds/moses/message"

const Type = "NOTICE"

// New creates a new NoticeMessage.
func New() message.Message {
	return message.New(Type)
}
