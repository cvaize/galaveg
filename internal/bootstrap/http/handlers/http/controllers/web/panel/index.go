package web

import (
	"galaveg/internal/bootstrap/http/context"
	"galaveg/internal/modules/view/layouts/home"
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
	//user := session.Get(userkey)
	//c.JSON(http.StatusOK, gin.H{"user": user})
	d, err := home.New(c, ctr.ctx.Services.App, ctr.ctx.Services.Locales, ctr.ctx.Services.Translator, session, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, home.TEMPLATE, d)
}
