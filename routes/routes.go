package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Http(router *gin.Engine) {
	apiRegister(router)
	webRegister(router)
	staticFilesRegister(router)
}

func Chat() {
	fmt.Println("Chat routes")
}
