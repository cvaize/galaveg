package routes

import (
	"galaveg/app/controllers/api/v1"
	"github.com/gin-gonic/gin"
)

func webRegister(router *gin.Engine) {
	var ctrl v1.Controller
	g := router.Group("/")
	g.GET("/", ctrl.Index)
}
