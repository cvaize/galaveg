package auth

import (
	"galaveg/internal/bootstrap/http/context"
	"galaveg/internal/modules/auth/cases/reset_password"
)

type Controller struct {
	ctx                  *context.Context
	resetPasswordContext *reset_password.Context
}

func NewController(ctx *context.Context) Controller {
	rl := ctx.Services.RateLimit
	as := ctx.Services.App
	auth := ctx.Services.Auth
	us := ctx.Services.Users
	ns := ctx.Services.Notifications
	kv := ctx.Infra.KV
	resetPasswordContext := reset_password.NewContext(as, auth, us, ns, kv, rl)
	return Controller{ctx, resetPasswordContext}
}
