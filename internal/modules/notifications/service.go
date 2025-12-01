package notifications

import (
	"fmt"
	"galaveg/internal/config"
	"galaveg/internal/modules/app"
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail"
	"galaveg/internal/modules/template"
	"galaveg/internal/modules/translator"
)

type Service = *ServiceImpl

type ServiceImpl struct {
	mailService  mail.Service
	buildContext *BuildNotificationContext
}

type BuildNotificationContext struct {
	Cfg               *config.Config
	TranslatorService translator.Service
	AppService        app.Service
	TemplateService   template.Service
}

func NewService(cfg *config.Config, mailService mail.Service, ts translator.Service, as app.Service, tmps template.Service) (*ServiceImpl, *errors.Error) {
	buildContext := &BuildNotificationContext{cfg, ts, as, tmps}
	return &ServiceImpl{mailService, buildContext}, nil
}

// NotificationsService The service is responsible for sending notifications via E-mail, SMS, Alerts, and Chat. Currently, only sending to E-mail is implemented.
func (s *ServiceImpl) Send(n Notification) []*errors.Error {
	if n.IsQueue(s.buildContext) {
		// TODO: Make a submission in the queue
	}
	var errs []*errors.Error
	channels := n.Via(s.buildContext)
	for _, channel := range channels {
		if channel == "mail" || channel == "email" {
			message, e := n.BuildEmailMessage(s.buildContext)
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
			errors.E500(e, "notifications.ServiceImpl.Send.Via", "")
		}
	}
	return errs
}
