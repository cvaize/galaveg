package auth

import (
	sessionModule "galaveg/internal/modules/sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctr *Controller) Logout(c *gin.Context) {
	session := sessions.Default(c)
	if sessionModule.ExistsUserId(ctr.ctx.Cfg, session) {
		session.Clear()
	}
	c.Redirect(http.StatusFound, "/")
}
