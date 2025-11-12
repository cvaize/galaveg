package routes

import (
	"galaveg/app/controllers/web"
	"galaveg/app/controllers/web/auth"
	"galaveg/app/controllers/web/users"
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
)

func webRegister(r *gin.Engine) {
	webCtrl := web.MustNewWebController(singleton.AS, singleton.TS)
	r.GET("/", webCtrl.Home)
	r.GET("/login", auth.Login)

	g := r.Group("/users")
	g.GET("/", users.Index)
}
