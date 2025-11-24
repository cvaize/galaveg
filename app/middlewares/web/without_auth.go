package web

import (
	"galaveg/bootstrap/providers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WithoutAuth(ctx *providers.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ctx.SS.ExistsUserId(sessions.Default(c)) {
			c.Redirect(http.StatusFound, "/panel")
			return
		}

		c.Next()
	}
}
