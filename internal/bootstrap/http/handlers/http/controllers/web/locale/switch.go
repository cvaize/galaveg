package locale

import (
	"galaveg/internal/bootstrap/http/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	ctx *context.Context
}

func NewController(ctx *context.Context) Controller {
	return Controller{ctx}
}

func (ctr *Controller) Switch(c *gin.Context) {
	v := c.PostForm("locale")

	if v == "" {
		v = ctr.ctx.Cfg.App.Locale
	}

	lenLocale := len(v)
	if lenLocale < 1 || lenLocale > 6 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	k := ctr.ctx.Cfg.App.LocaleCookieKey
	c.SetCookie(k, v, 0, "/", "", false, true)

	location := c.GetHeader("Referer")
	if location == "" {
		location = c.GetHeader("Origin")
	}
	if location == "" {
		location = "/"
	}
	c.Redirect(http.StatusFound, location)
}
