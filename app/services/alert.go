package services

import (
	"errors"
	"galaveg/app/dto"
	"github.com/gin-contrib/sessions"
)

const alertsKey = "alerts"

type AlertService struct {
	ES *ErrorService
}

func (s *AlertService) Flashes(session sessions.Session) ([]dto.Alert, *dto.Error) {
	var alerts []dto.Alert
	flashes := session.Flashes(alertsKey)
	if e := session.Save(); e != nil {
		return alerts, s.ES.E500(e, "AlertService.Flashes.FailedToSaveSession", "")
	}

	for _, flash := range flashes {
		alert, ok := flash.(dto.Alert)
		if !ok {
			return alerts, s.ES.E500(errors.New(""), "AlertService.Flashes.InvalidDeserializeAlert", "")
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

func (s *AlertService) AddFlash(session sessions.Session, alerts []dto.Alert) *dto.Error {
	for _, alert := range alerts {
		session.AddFlash(alert, alertsKey)
	}
	if e := session.Save(); e != nil {
		return s.ES.E500(e, "AlertService.AddFlash.FailedToSaveSession", "")
	}

	return nil
}
