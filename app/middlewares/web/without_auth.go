package web

import (
	"galaveg/bootstrap/providers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WithoutAuth(ctx *providers.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if user := session.Get(ctx.C.Session.StoreUserKey); user != nil {
			c.Redirect(http.StatusFound, "/panel")
			return
		}

		c.Next()
	}
}
