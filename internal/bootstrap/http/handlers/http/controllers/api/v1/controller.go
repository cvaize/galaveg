package v1

import (
	"galaveg/internal/bootstrap/http/context"
	dbModule "galaveg/internal/modules/db"
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
	search := "%ADMIN%"
	values := make([]interface{}, 2)
	values[0] = search
	values[1] = search
	filters := &dbModule.DbRepoFilters{[]string{"(name like ? or description like ?)"}, values}
	query := &dbModule.DbRepoQuery{
		Filters: filters,
	}
	dto, err := ctr.ctx.Services.Roles.All(query)
	debug.Dump(dto)
	debug.Dump(err)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
