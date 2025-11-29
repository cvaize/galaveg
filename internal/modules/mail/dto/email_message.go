package dto

type EmailMessage struct {
	Envelope    *Envelope
	Content     *Content
	Headers     *Headers
	Attachments []Attachment
}

func (s *EmailMessage) getEnvelope() *Envelope {
	return s.Envelope
}

func (s *EmailMessage) getContent() *Content {
	return s.Content
}

func (s *EmailMessage) getHeaders() *Headers {
	return s.Headers
}

func (s *EmailMessage) getAttachments() []Attachment {
	return s.Attachments
}
