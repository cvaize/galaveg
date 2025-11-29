package users

import (
	"galaveg/internal/bootstrap/http/context"
	"galaveg/internal/modules/view/layouts/list"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	ctx *context.Context
}

func NewController(ctx *context.Context) Controller {
	return Controller{ctx}
}

func (ctr *Controller) Index(c *gin.Context) {
	session := sessions.Default(c)
	d, err := list.New(c, ctr.ctx.Services.App, ctr.ctx.Services.Locales, ctr.ctx.Services.Translator, session, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, list.TEMPLATE, d)
}
