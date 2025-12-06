package auth

import (
	"galaveg/internal/modules/alerts"
	"galaveg/internal/modules/auth"
	"galaveg/internal/modules/errors"
	sessionsModule "galaveg/internal/modules/sessions"
	view "galaveg/internal/modules/view/layouts/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" binding:"required,min=6"`
}

func (ctr *Controller) Register(c *gin.Context) {
	session := sessions.Default(c)
	if sessionsModule.ExistsUserId(ctr.ctx.Cfg, session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}
	authService := ctr.ctx.Services.Auth
	as := ctr.ctx.Services.App
	ts := ctr.ctx.Services.Translator
	ls := ctr.ctx.Services.Locales

	locale := ls.Locale(c, nil)
	viewData := view.RegisterViewData{}
	reqData := RegisterRequest{}
	status := http.StatusOK
	if c.Request.Method == "POST" {
		valid := false

		if err := c.ShouldBind(&reqData); err != nil {
			errs := ts.TVE(locale, err)

			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(viewData.EmailErrors, e.GetMessage(ts.T(locale, "page.register.fields.email")))
				} else if e.Name == "Password" {
					viewData.PasswordErrors = append(viewData.PasswordErrors, e.GetMessage(ts.T(locale, "page.register.fields.password")))
				} else if e.Name == "ConfirmPassword" {
					viewData.ConfirmPasswordErrors = append(viewData.ConfirmPasswordErrors, e.GetMessage(ts.T(locale, "page.register.fields.confirm_password")))
				}
			}

		} else {
			if reqData.Password != reqData.ConfirmPassword {
				a := ts.T(locale, "page.register.fields.password")
				attributes := map[string]string{"attribute": a}
				viewData.PasswordErrors = append(viewData.PasswordErrors, "")
				viewData.ConfirmPasswordErrors = append(viewData.ConfirmPasswordErrors, ts.V(locale, "validation.confirmed", attributes))
			} else {
				valid = true
			}
		}

		if valid {
			e := authService.Register(auth.NewEmailVO(reqData.Email), auth.NewPasswordVO(reqData.Password))
			if e != nil {
				if e.Code == auth.ErrorRegisterDuplicateUser {
					status = http.StatusBadRequest
					viewData.Errors = append(viewData.Errors, ts.T(locale, auth.TranslateUserIsAlreadyRegistered))
				} else {
					status = e.Status
					viewData.Errors = append(viewData.Errors, ts.T(locale, errors.Translate500))
				}
			} else {
				alert := alerts.NewSuccessAlert(ts.T(locale, "alert.register.success"))
				//goland:noinspection GoUnhandledErrorResult
				alerts.AddFlash(session, []alerts.AlertDto{alert})
				c.Redirect(http.StatusFound, "/login")
				return
			}
		}

		viewData.EmailValue = reqData.Email
		viewData.PasswordValue = reqData.Password
		viewData.ConfirmPasswordValue = reqData.ConfirmPassword
	}

	d, err := view.NewRegister(c, as, ls, ts, session, &viewData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
