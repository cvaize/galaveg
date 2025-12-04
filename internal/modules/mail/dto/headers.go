package dto

type Headers struct {
	MessageId  string   // The message's message Id.
	References []string // The message IDs that are referenced by the message.
	Text       []string // The message's text headers.
}
