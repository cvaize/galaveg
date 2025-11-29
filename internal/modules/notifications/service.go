package notifications

import (
	"fmt"
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail"
)

type Service struct {
	mailService *mail.Service
}

func NewService(mailService *mail.Service) (*Service, *errors.Error) {
	return &Service{mailService}, nil
}

// NotificationsService The service is responsible for sending notifications via E-mail, SMS, Alerts, and Chat. Currently, only sending to E-mail is implemented.
func (s *Service) Send(n NotificationContract) []*errors.Error {
	if n.IsQueue() {
		// TODO: Make a submission in the queue
	}
	var errs []*errors.Error
	channels := n.Via()
	for _, channel := range channels {
		if channel == "mail" || channel == "email" {
			message, e := n.BuildEmailMessage()
			if e != nil {
				errs = append(errs, e)
			}
			e = s.mailService.Send(message)
			if e != nil {
				errs = append(errs, e)
			}
		} else {
			e := fmt.Errorf("notification Channel \"%s\" unsupported", channel)
			//goland:noinspection GoUnhandledErrorResult
			errors.E500(e, "notifications.Service.Send.Via", "")
		}
	}
	return errs
}
