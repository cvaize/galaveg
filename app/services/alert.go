package services

import (
	"errors"
	"galaveg/app/dto"
	"github.com/gin-contrib/sessions"
)

const alertsKey = "alerts"

type AlertService struct {
	es *ErrorService
}

func NewAlertService(es *ErrorService) (*AlertService, error) {
	return &AlertService{es}, nil
}

//goland:noinspection GoUnhandledErrorResult
func (s *AlertService) Flashes(session sessions.Session) []dto.Alert {
	var alerts []dto.Alert
	flashes := session.Flashes(alertsKey)
	if e := session.Save(); e != nil {
		s.es.E500(e, "AlertService.Flashes.FailedToSaveSession", "")
		return alerts
	}

	for _, flash := range flashes {
		alert, ok := flash.(dto.Alert)
		if !ok {
			s.es.E500(errors.New(""), "AlertService.Flashes.InvalidDeserializeAlert", "")
			return alerts
		}
		alerts = append(alerts, alert)
	}

	return alerts
}

func (s *AlertService) AddFlash(session sessions.Session, alerts []dto.Alert) *dto.Error {
	for _, alert := range alerts {
		session.AddFlash(alert, alertsKey)
	}
	if e := session.Save(); e != nil {
		return s.es.E500(e, "AlertService.AddFlash.FailedToSaveSession", "")
	}

	return nil
}
