package routes

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	apiRegister(router)
	webRegister(router)
}
