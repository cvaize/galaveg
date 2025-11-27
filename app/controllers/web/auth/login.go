package auth

import (
	view "galaveg/app/view/layouts/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

func (ctr *WebAuthController) Login(c *gin.Context) {
	session := sessions.Default(c)
	if ctr.ctx.S.SS.ExistsUserId(session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	locale := ctr.ctx.S.AS.Locale(c, nil)
	viewData := view.LoginViewData{}
	reqData := LoginRequest{}
	status := http.StatusOK

	if c.Request.Method == "POST" {
		if err := c.ShouldBind(&reqData); err != nil {
			errs := ctr.ctx.S.TS.TVE(locale, err)

			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(viewData.EmailErrors, e.GetMessage(ctr.ctx.S.TS.T(locale, "page.login.fields.email")))
				} else if e.Name == "Password" {
					viewData.PasswordErrors = append(viewData.PasswordErrors, e.GetMessage(ctr.ctx.S.TS.T(locale, "page.login.fields.password")))
				}
			}
		} else {

			userId, e := ctr.ctx.S.AuthS.Login(reqData.Email, reqData.Password)
			if e != nil {
				status = e.Status
				if status >= 500 {
					viewData.Errors = append(viewData.Errors, ctr.ctx.S.TS.T(locale, "error.500"))
				} else {
					if e.Code == "AuthService.Login.UserNotFound" {
						viewData.Errors = append(viewData.Errors, ctr.ctx.S.TS.T(locale, "error.AuthS.UserHasNotYetRegistered"))
					} else {
						viewData.Errors = append(viewData.Errors, ctr.ctx.S.TS.T(locale, "error.AuthS.CredentialsInvalid"))
					}
				}
			} else {
				e = ctr.ctx.S.SS.Login(session, userId)
				if e != nil {
					status = e.Status
					viewData.Errors = append(viewData.Errors, ctr.ctx.S.TS.T(locale, "error.500"))
				} else {
					c.Redirect(http.StatusFound, "/")
					return
				}
			}
		}

		viewData.EmailValue = reqData.Email
		viewData.PasswordValue = reqData.Password
	}

	d, err := view.NewLogin(c, ctr.ctx, session, &viewData)
	if err != nil {
		//goland:noinspection GoUnhandledErrorResult
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
