package routes

import (
	"galaveg/internal/bootstrap/http/context"
	v1 "galaveg/internal/bootstrap/http/handlers/http/controllers/api/v1"
	"github.com/gin-gonic/gin"
)

func apiRouter(router *gin.Engine, ctx *context.Context) {
	ctr := v1.NewController(ctx)
	g := router.Group("/api/v1")
	g.GET("/", ctr.Index)
}
