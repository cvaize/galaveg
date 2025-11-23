package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"github.com/gin-contrib/sessions"
)

type SessionService struct {
	cfg *config.Config
	ES  *ErrorService
}

func (s *SessionService) Login(session sessions.Session, userId dto.UserID) *dto.Error {
	session.Set(s.cfg.Session.StoreUserKey, userId)
	if e := session.Save(); e != nil {
		return s.ES.E500(e, "SessionService.Login.FailedToSaveSession", "")
	}
	return nil
}
