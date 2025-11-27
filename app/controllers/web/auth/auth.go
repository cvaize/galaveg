package auth

import "galaveg/bootstrap/providers"

type WebAuthController struct {
	ctx *providers.Context
}

func NewController(ctx *providers.Context) WebAuthController {
	return WebAuthController{ctx}
}
