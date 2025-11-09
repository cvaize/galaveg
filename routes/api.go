package routes

import (
	"galaveg/app/controllers/api/v1"
	"github.com/gin-gonic/gin"
)

func apiRegister(router *gin.Engine) {
	g := router.Group("/api/v1")
	g.GET("/", v1.Index)
}
