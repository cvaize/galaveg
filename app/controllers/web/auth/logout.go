package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctr *Controller) Logout(c *gin.Context) {
	session := ctr.ctx.SS.Default(c)
	if ctr.ctx.SS.ExistsUserId(session) {
		ctr.ctx.SS.Clear(session)
	}
	c.Redirect(http.StatusFound, "/")
}
