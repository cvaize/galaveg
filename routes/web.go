package routes

import (
	"galaveg/app/controllers/api/v1"
	"galaveg/app/controllers/web/auth"
	"github.com/gin-gonic/gin"
)

func webRegister(router *gin.Engine) {
	var ctrl v1.Controller
	g := router.Group("/")
	g.GET("/", ctrl.Index)
	g.GET("/login", auth.Login)
}
