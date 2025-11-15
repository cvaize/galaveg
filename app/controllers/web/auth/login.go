package auth

import (
	"galaveg/app/view/layouts/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	d, err := auth.NewLogin(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, auth.TEMPLATE, d)
}
