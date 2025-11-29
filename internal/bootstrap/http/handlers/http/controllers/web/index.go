package web

import (
	"galaveg/internal/bootstrap/http/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	ctx *context.Context
}

func NewController(ctx *context.Context) Controller {
	return Controller{ctx}
}

func (ctr *Controller) Index(c *gin.Context) {
	c.Redirect(http.StatusFound, "/panel")
	//session := sessions.Default(c)
	//user := session.Get(userkey)
	//c.JSON(http.StatusOK, gin.H{"user": user})
	//d, err := home.New(c, ctr.ctx, nil)
	//if err != nil {
	//	c.AbortWithError(http.StatusInternalServerError, err)
	//	return
	//}
	//
	//c.HTML(http.StatusOK, home.TEMPLATE, d)
}
