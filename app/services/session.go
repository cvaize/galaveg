package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"galaveg/utils/logger"
	"github.com/gin-contrib/sessions"
)

type SessionService struct {
	cfg *config.Config
	es  *ErrorService
}

func NewSessionService(cfg *config.Config, es *ErrorService) (*SessionService, error) {
	return &SessionService{cfg, es}, nil
}

func MustSessionService(cfg *config.Config, es *ErrorService) *SessionService {
	s, e := NewSessionService(cfg, es)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *SessionService) ExistsUserId(session sessions.Session) bool {
	user := session.Get(s.cfg.Session.StoreUserKey)
	return user != nil
}

func (s *SessionService) GetUserId(session sessions.Session) (dto.UserID, bool) {
	user := session.Get(s.cfg.Session.StoreUserKey)

	if user == nil {
		return 0, false
	}

	userId, ok := user.(dto.UserID)

	if !ok {
		logger.Errorf("(500) SessionService.GetUserId: %v", user)
	}

	return userId, ok
}

func (s *SessionService) Login(session sessions.Session, userId dto.UserID) *dto.Error {
	session.Set(s.cfg.Session.StoreUserKey, userId)
	if e := session.Save(); e != nil {
		return s.es.E500(e, "SessionService.Login.FailedToSaveSession", "")
	}
	return nil
}
