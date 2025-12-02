package reset_password

import (
	"context"
	"fmt"
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/users"
	"github.com/samber/lo"
	"time"
)

var (
	kvContext            = context.Background()
	codeLen              = 16
	codeCharset          = lo.AlphanumericCharset
	codeTTL              = 5 * time.Minute
	codeValue            = "1"
	rateLimitMaxAttempts = 1
	rateLimitTTL         = time.Minute
)

// TODO: RateLimit in Controller

func Send(ctx *Context, locale, email string) *errors.Error {
	user, e := ctx.us.FirstByEmail(email)
	if e != nil {
		return errors.E500(e, "auth.reset_password.FailedToGetUser", "")
	}

	if user == nil {
		return errors.E404(e, "auth.reset_password.UserNotFound", "")
	}

	rlKey := makeRLSendKey(user.ID)
	executed, e := ctx.rl.Attempt(rlKey, rateLimitMaxAttempts, rateLimitTTL)
	if e != nil {
		return errors.E500(e, "auth.reset_password.RateLimit.Attempt", "")
	}

	if !executed {
		message, e := ctx.rl.TtlMessage(locale, rlKey)
		if e != nil {
			return errors.E500(e, "auth.reset_password.RateLimit.TtlMessage", "")
		}
		return errors.E400(e, "auth.reset_password.RateLimit", message)
	}

	code := makeCode()
	key := makeKey(user.ID, code)

	// TODO: Изменить хранение на массив, иначе ключом могут сломать хранилище
	// Лучше всего наверное реализовать как это сделал Steam или шифровать код

	err := ctx.kv.SetEx(kvContext, key, codeValue, codeTTL).Err()
	if err != nil {
		return errors.E500(e, "auth.reset_password.FailedToSaveCode", "")
	}

	link := makeLink(ctx, email, code)

	notification := NewNotification(locale, email, link)

	errs := ctx.ns.Send(notification)

	if len(errs) > 0 {
		return errors.E500(e, "auth.reset_password.FailedToSendNotification", "")
	}

	return nil
}

func makeKey(id users.ID, code string) string {
	return fmt.Sprintf("auth.reset_password_link.code.%d_%s", id, code)
}

func makeRLSendKey(id users.ID) string {
	return fmt.Sprintf("auth.reset_password_link.%d", id)
}

func makeCode() string {
	return lo.RandomString(codeLen, codeCharset)
}

func makeLink(ctx *Context, email, code string) string {
	refUrl := ctx.as.RefUrl()
	refUrl = refUrl.JoinPath("reset-password-confirm")

	q := refUrl.Query()
	q.Set("code", code)
	q.Set("email", email)
	refUrl.RawQuery = q.Encode()

	return refUrl.String()
}
