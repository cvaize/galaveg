package notifications

import (
	"fmt"
	"galaveg/app/contracts"
	"galaveg/app/dto"
	"galaveg/bootstrap/providers"
)

// NotificationsService The service is responsible for sending notifications via E-mail, SMS, Alerts, and Chat. Currently, only sending to E-mail is implemented.
func Send(ctx *providers.Context, n contracts.NotificationContract) []*dto.Error {
	if n.IsQueue(ctx) {
		// TODO: Make a submission in the queue
	}
	var errs []*dto.Error
	channels := n.Via(ctx)
	for _, channel := range channels {
		if channel == "mail" || channel == "email" {
			message, e := n.BuildEmailMessage(ctx)
			if e != nil {
				errs = append(errs, e)
			}
			e = ctx.S.MS.Send(message)
			if e != nil {
				errs = append(errs, e)
			}
		} else {
			e := fmt.Errorf("notification Channel \"%s\" unsupported", channel)
			//goland:noinspection GoUnhandledErrorResult
			ctx.S.ES.E500(e, "NotificationsService.Send.Via", "")
		}
	}
	return errs
}
