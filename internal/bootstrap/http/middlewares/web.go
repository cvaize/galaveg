package middlewares

import (
	"galaveg/config"
	actions "galaveg/internal/modules/sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// WebAuthRequired is a middleware that checks if the user has a valid session.
// It should be used on routes that require authentication.
// If no valid session exists, it aborts the request with 401 Unauthorized.
func WebAuthRequired(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session from the request context
		session := sessions.Default(c)

		// Try to get the user from the session
		if user := session.Get(cfg.Session.StoreUserKey); user == nil {
			// No user in session, abort the request
			c.Redirect(http.StatusFound, "/login")
			return
		}

		// User is authenticated, continue to the next handler
		c.Next()
	}
}

func WebWithoutAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if actions.ExistsUserId(cfg, sessions.Default(c)) {
			c.Redirect(http.StatusFound, "/panel")
			return
		}

		c.Next()
	}
}
