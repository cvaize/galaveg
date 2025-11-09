package routes

import (
	"galaveg/app/controllers/web"
	"galaveg/app/controllers/web/auth"
	"galaveg/app/controllers/web/users"
	"github.com/gin-gonic/gin"
)

func webRegister(r *gin.Engine) {
	r.GET("/", web.Index)
	r.GET("/login", auth.Login)

	g := r.Group("/users")
	g.GET("/", users.Index)
}
