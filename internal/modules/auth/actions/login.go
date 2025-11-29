package actions

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/hash"
	"galaveg/internal/modules/users"
)

func Login(us *users.Service, hs *hash.Service, email, password string) (users.UserID, *errors.Error) {
	user, e := us.FirstByEmail(email)
	if e != nil {
		// Failed to get user
		return 0, errors.E500(e, "auth.actions.Login.FailedToGetUser", "")
	}

	if user == nil {
		// User not found
		return 0, errors.E404(e, "auth.actions.Login.UserNotFound", "")
	}

	is, e := hs.VerifyPassword(password, user.Password)
	if e != nil {
		// Failed to verify password
		return 0, errors.E500(e, "auth.actions.Login.FailedToVerifyPassword", "")
	}
	if !is {
		// The password does not match the saved hash
		return 0, errors.E401(e, "auth.actions.Login.Unauthorized", "")
	}

	return user.ID, nil
}
