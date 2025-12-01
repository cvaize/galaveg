package actions

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/hash"
	"galaveg/internal/modules/users"
	"strings"
)

func Register(us users.Service, hs hash.Service, email, password string) (users.UserID, *errors.Error) {
	exists, e := us.ExistsByEmail(email)
	if e != nil {
		// Failed to check for user existence
		return 0, errors.E500(e, "auth.actions.Register.FailedToGetUser", "")
	}
	if exists {
		// This user is already registered
		return 0, errors.E400(e, "auth.actions.Register.DuplicateUser", "")
	}

	password, e = hs.HashPassword(password)
	if e != nil {
		// Failed to generate password hash
		return 0, errors.E500(e, "auth.actions.Register.HashPasswordFail", "")
	}

	userId, err := us.Create(&users.User{Email: email, Password: password})
	if err != nil {
		eStr := err.Error()
		if strings.Contains(eStr, "Duplicate entry") {
			if strings.Contains(eStr, ".email'") {
				// There is already such a user
				return 0, errors.E400(e, "auth.actions.Register.DuplicateUser", "")
			}
			// There is already such a user, not intended behavior
			return 0, errors.E500(e, "auth.actions.Register.DuplicateUser", "")
		}
		// Failed to register user
		return 0, errors.E500(e, "auth.actions.Register.InsertNewUserFail", "")
	}

	return userId, nil
}
