package mailables

import "net/mail"

type Envelope struct {
	From     *mail.Address     // The address sending the message.
	To       []*mail.Address   // The recipients of the message.
	Cc       []*mail.Address   // The recipients receiving a copy of the message.
	Bcc      []*mail.Address   // The recipients receiving a blind copy of the message.
	ReplyTo  []*mail.Address   // The recipients that should be replied to.
	Subject  string            // The subject of the message.
	Tags     []string          // The message's tags.
	Metadata map[string]string // The message's meta data.
}
