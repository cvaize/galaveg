package web

import (
	"galaveg/app/view/layouts/home"
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
	//session := sessions.Default(c)
	//user := session.Get(userkey)
	//c.JSON(http.StatusOK, gin.H{"user": user})
	d, err := home.New(c, ctr.ctx, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, home.TEMPLATE, d)
}
