package users

import (
	"galaveg/app/view/layouts/list"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	ctx *providers.Context
}

func NewController(ctx *providers.Context) Controller {
	return Controller{ctx}
}

func (ctr *Controller) Index(c *gin.Context) {
	d, err := list.New(c, ctr.ctx, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, list.TEMPLATE, d)
}
