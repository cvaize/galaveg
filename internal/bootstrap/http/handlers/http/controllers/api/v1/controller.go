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
	search := "%ADMIN%"
	values := make([]interface{}, 2)
	values[0] = search
	values[1] = search
	whereClauses := []string{"(name like ? or description like ?)"}
	columns := []string{"id", "name"}
	orderBy := "name ASC"
	records, totalRecords, totalPages, err := ctr.ctx.Services.Roles.Paginate(1, 10, values, whereClauses, columns, orderBy)
	debug.Dump(records)
	debug.Dump(totalRecords)
	debug.Dump(totalPages)
	debug.Dump(err)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
