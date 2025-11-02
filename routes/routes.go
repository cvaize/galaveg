package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Http(router *gin.Engine) {
	apiRegister(router)
	webRegister(router)
}

func Chat() {
	fmt.Println("Chat routes")
}
