package web

import (
	"galaveg/app/view/layouts/home"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	d, err := home.New(c, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, home.TEMPLATE, d)
}
