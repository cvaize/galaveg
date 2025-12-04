package v1

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
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
