package sessions

import (
	"galaveg/internal/config"
	"galaveg/internal/modules/errors"
	"github.com/gin-contrib/sessions"
)

func ExistsUserId(cfg *config.Config, session sessions.Session) bool {
	user := session.Get(cfg.Session.StoreUserKey)
	return user != nil
}

func GetUserId(cfg *config.Config, session sessions.Session) (int64, bool) {
	user := session.Get(cfg.Session.StoreUserKey)

	if user == nil {
		return 0, false
	}

	userId, ok := user.(int64)

	if !ok {
		//goland:noinspection GoUnhandledErrorResult
		errors.E500(nil, "sessions.actions.GetUserId.Id", "")
	}

	return userId, ok
}

func Login(cfg *config.Config, session sessions.Session, userId int64) *errors.Error {
	session.Set(cfg.Session.StoreUserKey, userId)
	if e := session.Save(); e != nil {
		return errors.E500(e, "sessions.actions.Login.FailedToSaveSession", "")
	}
	return nil
}
