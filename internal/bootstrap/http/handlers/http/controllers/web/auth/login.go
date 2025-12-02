package auth

import (
	"galaveg/internal/modules/auth"
	sessionsModule "galaveg/internal/modules/sessions"
	view "galaveg/internal/modules/view/layouts/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

func (ctr *Controller) Login(c *gin.Context) {
	session := sessions.Default(c)
	if sessionsModule.ExistsUserId(ctr.ctx.Cfg, session) {
		//if ctr.ctx.S.SS.ExistsUserId(session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}
	authService := ctr.ctx.Services.Auth
	as := ctr.ctx.Services.App
	ts := ctr.ctx.Services.Translator
	ls := ctr.ctx.Services.Locales

	locale := ctr.ctx.Services.Locales.Locale(c, nil)
	viewData := view.LoginViewData{}
	reqData := LoginRequest{}
	status := http.StatusOK

	if c.Request.Method == "POST" {
		if err := c.ShouldBind(&reqData); err != nil {
			errs := ts.TVE(locale, err)

			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(viewData.EmailErrors, e.GetMessage(ts.T(locale, "page.login.fields.email")))
				} else if e.Name == "Password" {
					viewData.PasswordErrors = append(viewData.PasswordErrors, e.GetMessage(ts.T(locale, "page.login.fields.password")))
				}
			}
		} else {
			userId, e := authService.Login(auth.NewEmailVO(reqData.Email), auth.NewPasswordVO(reqData.Password))
			if e != nil {
				status = e.Status
				if status >= 500 {
					viewData.Errors = append(viewData.Errors, ts.T(locale, "error.500"))
				} else {
					if e.Code == "AuthService.Login.UserNotFound" {
						viewData.Errors = append(viewData.Errors, ts.T(locale, "error.AuthS.UserHasNotYetRegistered"))
					} else {
						viewData.Errors = append(viewData.Errors, ts.T(locale, "error.AuthS.CredentialsInvalid"))
					}
				}
			} else {
				e = sessionsModule.Login(ctr.ctx.Cfg, session, userId)
				if e != nil {
					status = e.Status
					viewData.Errors = append(viewData.Errors, ts.T(locale, "error.500"))
				} else {
					c.Redirect(http.StatusFound, "/")
					return
				}
			}
		}

		viewData.EmailValue = reqData.Email
		viewData.PasswordValue = reqData.Password
	}

	d, err := view.NewLogin(c, as, ls, ts, session, &viewData)
	if err != nil {
		//goland:noinspection GoUnhandledErrorResult
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
