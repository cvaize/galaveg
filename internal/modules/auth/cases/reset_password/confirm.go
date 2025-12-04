package reset_password

import (
	"galaveg/internal/modules/errors"
	"galaveg/pkg/debug"
)

// TODO: RateLimit in Controller

func Confirm(ctx *Context, code, email, password string) *errors.Error {
	user, e := ctx.us.FirstByEmail(email)
	if e != nil {
		return errors.E500(e, "auth.reset_password.FailedToGetUser", "")
	}

	if user == nil {
		return errors.E404(e, "auth.reset_password.UserNotFound", "")
	}

	key := makeKey(user.Id, code)

	val, err := ctx.kv.GetDel(kvContext, key).Result()
	if err != nil {
		debug.Dump(err)
		// TODO: Если 404, то тоже нужно отправлять "auth.reset_password.CodeIsNotEqual"
		return errors.E500(e, "auth.reset_password.FailedToGetCode", "")
	}

	if val != codeValue {
		return errors.E400(e, "auth.reset_password.CodeIsNotEqual", "")
	}

	err = ctx.auth.UpdatePassword(user.Id, password)
	if err != nil {
		return errors.E500(e, "auth.reset_password.UpdatePassword", "")
	}

	return nil
}
