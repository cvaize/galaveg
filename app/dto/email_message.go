package dto

import "galaveg/app/dto/mailables"

type EmailMessage struct {
	Envelope    *mailables.Envelope
	Content     *mailables.Content
	Headers     *mailables.Headers
	Attachments []mailables.Attachment
}
