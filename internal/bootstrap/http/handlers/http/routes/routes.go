package routes

import (
	"galaveg/internal/bootstrap/http/context"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, ctx *context.Context) {
	apiRouter(router, ctx)
	webRouter(router, ctx)
	staticFilesRouter(router, ctx)
}
