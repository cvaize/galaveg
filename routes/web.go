package routes

import (
	"galaveg/app/controllers/web"
	"galaveg/app/controllers/web/auth"
	"galaveg/app/controllers/web/locale"
	"galaveg/app/controllers/web/users"
	"github.com/gin-gonic/gin"
)

func webRegister(r *gin.Engine) {
	r.GET("/", web.Home)
	r.GET("/login", auth.Login)
	r.POST("/locale/switch", locale.Switch)
	r.POST("/logout", auth.Logout)
	r.GET("/register", auth.Register)
	r.POST("/register", auth.Register)
	r.GET("/reset-password", auth.ResetPassword)
	r.POST("/reset-password", auth.ResetPassword)
	r.GET("/reset-password-confirm", auth.ResetPasswordConfirm)
	r.POST("/reset-password-confirm", auth.ResetPasswordConfirm)

	g := r.Group("/users")
	g.GET("/", users.Index)
}
