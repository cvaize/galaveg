package users

import (
	"galaveg/app/view/layouts/list"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	d, err := list.New(c, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, list.TEMPLATE, d)
}
