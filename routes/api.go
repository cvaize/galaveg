package routes

import (
	"galaveg/app/controllers/api/v1"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
)

func apiRegister(router *gin.Engine, ctx *providers.Context) {
	g := router.Group("/api/v1")
	g.GET("/", v1.Index)
}
