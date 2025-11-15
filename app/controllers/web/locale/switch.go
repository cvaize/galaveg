package locale

import (
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Switch(c *gin.Context) {
	v := c.PostForm("locale")

	if v == "" {
		v = singleton.C.App.Locale
	}

	lenLocale := len(v)
	if lenLocale < 1 || lenLocale > 6 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	k := singleton.C.App.LocaleCookieKey
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
