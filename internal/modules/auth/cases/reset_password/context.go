package reset_password

import (
	"galaveg/internal/infrastructures/kv"
	"galaveg/internal/modules/app"
	"galaveg/internal/modules/auth"
	"galaveg/internal/modules/notifications"
	"galaveg/internal/modules/rate_limit"
	"galaveg/internal/modules/users"
)

type Context struct {
	as   *app.Service
	auth *auth.Service
	us   *users.Service
	ns   *notifications.Service
	rl   *rate_limit.Service
	kv   kv.KV
}

func NewContext(as *app.Service, auth *auth.Service, us *users.Service, ns *notifications.Service, kv kv.KV, rl *rate_limit.Service) *Context {
	return &Context{as, auth, us, ns, rl, kv}
}
