package mail

type Envelope struct {
	// The address sending the message.
	From Address
	// The recipients of the message.
	To []Address
	// The recipients receiving a copy of the message.
	Cc []Address
	// The recipients receiving a blind copy of the message.
	Bcc []Address
	// The recipients that should be replied to.
	ReplyTo []Address
	// The subject of the message.
	Subject string
	// The message's tags.
	Tags     []string
	Metadata map[string]string
	//HtmlBody string
	//TextBody string
}

//
///**
// * The message's meta data.
// *
// * @var array
// */
//public $metadata = [];
//
///**
// * The message's Symfony Message customization callbacks.
// *
// * @var array
// */
//public $using = [];
