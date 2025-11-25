package mailables

type Envelope struct {
	From     Address           // The address sending the message.
	To       []Address         // The recipients of the message.
	Cc       []Address         // The recipients receiving a copy of the message.
	Bcc      []Address         // The recipients receiving a blind copy of the message.
	ReplyTo  []Address         // The recipients that should be replied to.
	Subject  string            // The subject of the message.
	Tags     []string          // The message's tags.
	Metadata map[string]string // The message's meta data.
}
