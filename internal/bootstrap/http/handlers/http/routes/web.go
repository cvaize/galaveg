package routes

import (
	"galaveg/internal/bootstrap/http/context"
	"galaveg/internal/bootstrap/http/handlers/http/controllers/web"
	"galaveg/internal/bootstrap/http/handlers/http/controllers/web/auth"
	"galaveg/internal/bootstrap/http/handlers/http/controllers/web/locale"
	"galaveg/internal/bootstrap/http/handlers/http/controllers/web/panel/users"
	mWeb "galaveg/internal/bootstrap/http/handlers/http/middlewares/web"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func webRouter(router *gin.Engine, ctx *context.Context) {
	r := router.Group("/")
	r.Use(sessions.Sessions(ctx.Cfg.Session.CookieKey, ctx.Infra.SessionStore))
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
	p.Use(mWeb.AuthRequired(ctx.Cfg))
	p.GET("/", wCtr.Index)
	p.GET("/users", uCtr.Index)
}
