package routes

import (
	"galaveg/app/controllers/web"
	"galaveg/app/controllers/web/auth"
	"galaveg/app/controllers/web/locale"
	"galaveg/app/controllers/web/panel/users"
	mWeb "galaveg/app/middlewares/web"
	"galaveg/bootstrap/providers"
	"galaveg/utils"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func webRegister(router *gin.Engine, ctx *providers.Context) {
	store := utils.Must(providers.NewSessionStore(ctx.Cfg))

	r := router.Group("/")
	r.Use(sessions.Sessions(ctx.Cfg.Session.CookieKey, store))
	wCtr := web.NewController(ctx)
	r.GET("/", wCtr.Index)

	// Public
	aCtr := auth.NewController(ctx)
	lCtr := locale.NewController(ctx)
	r.GET("/login", aCtr.Login)
	r.POST("/login", aCtr.Login)
	r.POST("/locale/switch", lCtr.Switch)
	r.POST("/logout", aCtr.Logout)
	r.GET("/register", aCtr.Register)
	r.POST("/register", aCtr.Register)
	r.GET("/reset-password", aCtr.ResetPassword)
	r.POST("/reset-password", aCtr.ResetPassword)
	r.GET("/reset-password-confirm", aCtr.ResetPasswordConfirm)
	r.POST("/reset-password-confirm", aCtr.ResetPasswordConfirm)

	// Private
	uCtr := users.NewController(ctx)
	p := r.Group("/panel")
	p.Use(mWeb.AuthRequired(ctx))
	p.GET("/", wCtr.Index)
	p.GET("/users", uCtr.Index)
}
