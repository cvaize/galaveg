package auth

import "galaveg/bootstrap/providers"

type Controller struct {
	ctx *providers.Context
}

func NewController(ctx *providers.Context) Controller {
	return Controller{ctx}
}
