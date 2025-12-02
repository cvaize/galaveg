package auth

import (
	"galaveg/internal/modules/alerts"
	"galaveg/internal/modules/auth/cases/reset_password"
	sessionsModule "galaveg/internal/modules/sessions"
	view "galaveg/internal/modules/view/layouts/auth"
	"galaveg/pkg/debug"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPasswordRequest struct {
	Email string `form:"email" binding:"required,email"`
}

func (ctr *Controller) ResetPassword(c *gin.Context) {
	session := sessions.Default(c)

	if sessionsModule.ExistsUserId(ctr.ctx.Cfg, session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}
	as := ctr.ctx.Services.App
	ts := ctr.ctx.Services.Translator
	ls := ctr.ctx.Services.Locales

	locale := ls.Locale(c, nil)
	viewData := view.ResetPasswordViewData{}
	reqData := ResetPasswordRequest{}
	status := http.StatusOK

	if c.Request.Method == "POST" {
		if err := c.ShouldBind(&reqData); err != nil {

			errs := ts.TVE(locale, err)
			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(
						viewData.EmailErrors,
						e.GetMessage(ts.T(locale, "page.reset_password.fields.email")),
					)
				}
			}

		} else {
			e := reset_password.Send(ctr.resetPasswordContext, locale, reqData.Email)

			if e != nil {
				debug.Dump(e)
				viewData.Errors = append(
					viewData.Errors,
					ts.T(locale, "alert.reset_password.fail"),
				)
			} else {
				alert := alerts.NewSuccessAlert(ts.T(locale, "alert.reset_password.success"))
				//goland:noinspection GoUnhandledErrorResult
				alerts.AddFlash(session, []alerts.AlertDto{alert})

				c.Redirect(http.StatusFound, "/login")
				return
			}
		}

		viewData.EmailValue = reqData.Email
	}

	d, err := view.NewResetPassword(c, as, ls, ts, session, &viewData)
	if err != nil {
		//goland:noinspection GoUnhandledErrorResult
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
