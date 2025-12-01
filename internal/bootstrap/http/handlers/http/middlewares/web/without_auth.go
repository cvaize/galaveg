package web

import (
	"galaveg/internal/config"
	sessionsActions "galaveg/internal/modules/sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WithoutAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessionsActions.ExistsUserId(cfg, sessions.Default(c)) {
			c.Redirect(http.StatusFound, "/panel")
			return
		}

		c.Next()
	}
}
