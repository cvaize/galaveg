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
	//search := "%ADMIN%"
	//values := make([]interface{}, 2)
	//values[0] = search
	//values[1] = search
	//whereClauses := []string{"(name like ? or description like ?)"}
	//totalRecords, err := ctr.ctx.Services.Roles.Count(values, whereClauses)
	//debug.Dump(totalRecords)
	//debug.Dump(err)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
