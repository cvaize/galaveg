package routes

import (
	"galaveg/app/controllers/api/v1"
	"github.com/gin-gonic/gin"
)

func apiRegister(router *gin.Engine) {
	var ctrl v1.Controller
	g := router.Group("/api/v1")
	g.GET("/", ctrl.Index)
}
