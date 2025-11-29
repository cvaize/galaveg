package notifications

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail/dto"
)

type NotificationContract interface {
	IsQueue() bool
	// Via Get the notification's delivery channels. Example: []string{"mail"}
	Via() []string
	BuildEmailMessage() (*dto.EmailMessage, *errors.Error)
}
