package alerts

import (
	"galaveg/internal/modules/errors"
	"github.com/gin-contrib/sessions"
)

const alertsKey = "alerts"

func Flashes(session sessions.Session) []AlertDto {
	var alerts []AlertDto
	flashes := session.Flashes(alertsKey)
	if e := session.Save(); e != nil {
		//goland:noinspection GoUnhandledErrorResult
		errors.E500(e, "alerts.Flashes.FailedToSaveSession", "")
		return alerts
	}

	for _, flash := range flashes {
		alert, ok := flash.(AlertDto)
		if !ok {
			//goland:noinspection GoUnhandledErrorResult
			errors.E500(nil, "alerts.Flashes.InvalidDeserializeAlert", "")
			return alerts
		}
		alerts = append(alerts, alert)
	}

	return alerts
}

func AddFlash(session sessions.Session, alerts []AlertDto) *errors.Error {
	for _, alert := range alerts {
		session.AddFlash(alert, alertsKey)
	}
	if e := session.Save(); e != nil {
		return errors.E500(e, "alerts.AddFlash.FailedToSaveSession", "")
	}

	return nil
}
