package sessions

import (
	"galaveg/config"
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/users"
	"github.com/gin-contrib/sessions"
)

func ExistsUserId(cfg *config.Config, session sessions.Session) bool {
	user := session.Get(cfg.Session.StoreUserKey)
	return user != nil
}

func GetUserId(cfg *config.Config, session sessions.Session) (users.UserID, bool) {
	user := session.Get(cfg.Session.StoreUserKey)

	if user == nil {
		return 0, false
	}

	userId, ok := user.(users.UserID)

	if !ok {
		//goland:noinspection GoUnhandledErrorResult
		errors.E500(nil, "sessions.actions.GetUserId.UserID", "")
	}

	return userId, ok
}

func Login(cfg *config.Config, session sessions.Session, userId users.UserID) *errors.Error {
	session.Set(cfg.Session.StoreUserKey, userId)
	if e := session.Save(); e != nil {
		return errors.E500(e, "sessions.actions.Login.FailedToSaveSession", "")
	}
	return nil
}
