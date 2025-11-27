package auth

import (
	notifications2 "galaveg/app/actions/notifications"
	"galaveg/app/dto"
	"galaveg/app/notifications"
	view "galaveg/app/view/layouts/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPasswordRequest struct {
	Email string `form:"email" binding:"required,email"`
}

func (ctr *WebAuthController) ResetPassword(c *gin.Context) {
	session := sessions.Default(c)

	if ctr.ctx.S.SS.ExistsUserId(session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	locale := ctr.ctx.S.AS.Locale(c, nil)
	viewData := view.ResetPasswordViewData{}
	reqData := ResetPasswordRequest{}
	status := http.StatusOK

	if c.Request.Method == "POST" {
		if err := c.ShouldBind(&reqData); err != nil {

			errs := ctr.ctx.S.TS.TVE(locale, err)
			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(
						viewData.EmailErrors,
						e.GetMessage(ctr.ctx.S.TS.T(locale, "page.reset_password.fields.email")),
					)
				}
			}

		} else {
			sendErr := notifications2.Send(ctr.ctx, notifications.NewResetPassword(locale, reqData.Email))
			if sendErr != nil && len(sendErr) > 0 {
				viewData.Errors = append(
					viewData.Errors,
					ctr.ctx.S.TS.T(locale, "alert.reset_password.fail"),
				)

			} else {
				alert := dto.NewSuccessAlert(ctr.ctx.S.TS.T(locale, "alert.reset_password.success"))
				//goland:noinspection GoUnhandledErrorResult
				ctr.ctx.S.AlS.AddFlash(session, []dto.Alert{alert})

				c.Redirect(http.StatusFound, "/login")
				return
			}
		}

		viewData.EmailValue = reqData.Email
	}

	d, err := view.NewResetPassword(c, ctr.ctx, session, &viewData)
	if err != nil {
		//goland:noinspection GoUnhandledErrorResult
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
