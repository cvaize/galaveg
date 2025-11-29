package auth

import "galaveg/internal/bootstrap/http/context"

type Controller struct {
	ctx *context.Context
}

func NewController(ctx *context.Context) Controller {
	return Controller{ctx}
}
