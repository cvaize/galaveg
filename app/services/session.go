package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"galaveg/utils/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionService struct {
	cfg *config.Config
	ES  *ErrorService
}

func (s *SessionService) Default(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}

func (s *SessionService) ExistsUserId(session sessions.Session) bool {
	user := session.Get(s.cfg.Session.StoreUserKey)
	return user != nil
}

func (s *SessionService) Clear(session sessions.Session) {
	session.Clear()
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
		return s.ES.E500(e, "SessionService.Login.FailedToSaveSession", "")
	}
	return nil
}
