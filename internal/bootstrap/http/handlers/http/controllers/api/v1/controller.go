package v1

import (
	"galaveg/internal/bootstrap/http/context"
	"galaveg/pkg/debug"
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
	f := make(map[string]interface{})
	f["id"] = 2
	dto, err := ctr.ctx.Services.Roles.AllIds(f, "")
	debug.Dump(dto)
	debug.Dump(err)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
