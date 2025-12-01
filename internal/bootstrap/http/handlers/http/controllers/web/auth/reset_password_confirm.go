package auth

import (
	sessionsModule "galaveg/internal/modules/sessions"
	"galaveg/internal/modules/view/layouts/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPasswordConfirmRequest struct {
	Code            string `form:"code" binding:"required,min=6"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" binding:"required,min=6"`
}

func (ctr *Controller) ResetPasswordConfirm(c *gin.Context) {
	session := sessions.Default(c)
	if sessionsModule.ExistsUserId(ctr.ctx.Cfg, session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}
	as := ctr.ctx.Services.App
	ts := ctr.ctx.Services.Translator
	ls := ctr.ctx.Services.Locales

	d, err := auth.NewResetPasswordConfirm(c, as, ls, ts, session)
	if err != nil {
		//goland:noinspection GoUnhandledErrorResult
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, auth.TEMPLATE, d)
}
