package routes

import (
	"fmt"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
)

func Http(router *gin.Engine, ctx *providers.Context) {
	apiRegister(router, ctx)
	webRegister(router, ctx)
	staticFilesRegister(router, ctx)
}

func Chat() {
	fmt.Println("Chat routes")
}
