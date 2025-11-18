package routes

import (
	"galaveg/app/controllers/web"
	"galaveg/app/controllers/web/auth"
	"galaveg/app/controllers/web/locale"
	"galaveg/app/controllers/web/users"
	"galaveg/app/middlewares"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func webRegister(router *gin.Engine, ctx *providers.Context) {
	store, err := redis.NewStore(100, "tcp", ctx.C.Redis.Host, ctx.C.Redis.Username, ctx.C.Redis.Password, []byte(ctx.C.App.Key))
	if err != nil {
		panic(err)
	}

	r := router.Group("/")
	r.Use(sessions.Sessions(ctx.C.Auth.CookieKey, store))
	wCtr := web.NewController(ctx)
	r.GET("/", wCtr.Index)

	// Public
	aCtr := auth.NewController(ctx)
	lCtr := locale.NewController(ctx)
	r.GET("/login", aCtr.Login)
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
	p.Use(middlewares.AuthRequired(ctx))
	p.GET("/", wCtr.Index)
	p.GET("/users", uCtr.Index)
}
