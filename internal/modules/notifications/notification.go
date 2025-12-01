package notifications

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail/dto"
)

type Notification interface {
	IsQueue(ctx *BuildNotificationContext) bool
	// Via Get the notification's delivery channels. Example: []string{"mail"}
	Via(ctx *BuildNotificationContext) []string
	BuildEmailMessage(ctx *BuildNotificationContext) (*dto.EmailMessage, *errors.Error)
}
