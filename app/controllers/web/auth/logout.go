package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctr *WebAuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	if ctr.ctx.S.SS.ExistsUserId(session) {
		session.Clear()
	}
	c.Redirect(http.StatusFound, "/")
}
