package auth

import (
	"galaveg/app/dto"
	view "galaveg/app/view/layouts/auth"
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
			url := ctr.ctx.AS.CloneUrl()
			url.JoinPath("reset-password-confirm")
			q := url.Query()
			q.Add("code", "code123")
			q.Add("email", reqData.Email)

			to := reqData.Email
			subject := "subject 123"
			html := "html 123"
			txt := "txt 123"
			e := ctr.ctx.MS.SendSimpleEmailMessage(to, subject, html, txt)
			if e != nil {
				logger.Errorf("(500) Controller.ResetPassword.SendSimpleEmailMessage: %v", e)
				viewData.Errors = append(viewData.Errors, ctr.ctx.TS.T(locale, "alert.reset_password.fail"))
			} else {
				alert := dto.NewSuccessAlert(ctr.ctx.TS.T(locale, "alert.reset_password.success"))
				//goland:noinspection GoUnhandledErrorResult
				ctr.ctx.AlS.AddFlash(session, []dto.Alert{alert})
			}
		}
	}
	d, err := view.NewResetPassword(c, ctr.ctx, &viewData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
