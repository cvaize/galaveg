package middlewares

import (
	"galaveg/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApiAuthRequired is a middleware that checks if the user has a valid session.
// It should be used on routes that require authentication.
// If no valid session exists, it aborts the request with 401 Unauthorized.
func ApiAuthRequired(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session from the request context
		session := sessions.Default(c)

		// Try to get the user from the session
		if user := session.Get(cfg.Session.StoreUserKey); user == nil {
			// No user in session, abort the request
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// User is authenticated, continue to the next handler
		c.Next()
	}
}
