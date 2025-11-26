package auth

import (
	"bytes"
	"galaveg/app/dto"
	view "galaveg/app/view/layouts/auth"
	auth "galaveg/app/view/layouts/email/reset_password"
	"galaveg/utils/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPasswordRequest struct {
	Email string `form:"email" binding:"required,email"`
}

func (ctr *Controller) ResetPassword(c *gin.Context) {
	session := sessions.Default(c)
	if ctr.ctx.SS.ExistsUserId(session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	locale := ctr.ctx.AS.Locale(c, nil)
	viewData := view.ResetPasswordViewData{}
	reqData := ResetPasswordRequest{}
	status := http.StatusOK

	if c.Request.Method == "POST" {
		if err := c.ShouldBind(&reqData); err != nil {
			errs := ctr.ctx.TS.TVE(locale, err)

			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(viewData.EmailErrors, e.GetMessage(ctr.ctx.TS.T(locale, "page.reset_password.fields.email")))
				}
			}
		} else {

			refUrl := ctr.ctx.AS.RefUrl()
			refUrl = refUrl.JoinPath("reset-password-confirm")
			q := refUrl.Query()
			q.Set("code", "code123")
			q.Set("email", reqData.Email)
			refUrl.RawQuery = q.Encode()
			url := refUrl.String()

			emailViewData := auth.ViewData{ResetPasswordLink: url}
			d, err1 := auth.New(c, ctr.ctx, &emailViewData)
			if err1 != nil {
				c.AbortWithError(http.StatusInternalServerError, err1)
				return
			}
			var tpl bytes.Buffer
			if templateError := ctr.ctx.Html.ExecuteTemplate(&tpl, auth.TEMPLATE, d); templateError != nil {
				logger.Errorf("Error: %v", templateError)
			} else {
				to := reqData.Email
				subject := ctr.ctx.TS.T(locale, "mail.reset_password.subject")
				html := tpl.String()
				txt := url

				// TODO: Separate rendering of E-mail templates into a separate entity
				e := ctr.ctx.MS.SyncSendSimpleEmail(to, subject, html, txt)
				if e != nil {
					logger.Errorf("(500) Controller.ResetPassword.SyncSendSimpleEmail: %v", e)
					viewData.Errors = append(viewData.Errors, ctr.ctx.TS.T(locale, "alert.reset_password.fail"))
				} else {
					alert := dto.NewSuccessAlert(ctr.ctx.TS.T(locale, "alert.reset_password.success"))
					//goland:noinspection GoUnhandledErrorResult
					ctr.ctx.AlS.AddFlash(session, []dto.Alert{alert})
				}
			}
		}
		viewData.EmailValue = reqData.Email
	}
	d, err := view.NewResetPassword(c, ctr.ctx, &viewData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
