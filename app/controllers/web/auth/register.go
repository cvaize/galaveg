package auth

import (
	"galaveg/app/dto"
	view "galaveg/app/view/layouts/auth"
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
	if ctr.ctx.SS.ExistsUserId(session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	locale := ctr.ctx.AS.Locale(c, nil)
	viewData := view.RegisterViewData{}
	reqData := RegisterRequest{}
	status := http.StatusOK
	if c.Request.Method == "POST" {
		valid := false

		if err := c.ShouldBind(&reqData); err != nil {
			errs := ctr.ctx.TS.TVE(locale, err)

			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(viewData.EmailErrors, e.GetMessage(ctr.ctx.TS.T(locale, "page.register.fields.email")))
				} else if e.Name == "Password" {
					viewData.PasswordErrors = append(viewData.PasswordErrors, e.GetMessage(ctr.ctx.TS.T(locale, "page.register.fields.password")))
				} else if e.Name == "ConfirmPassword" {
					viewData.ConfirmPasswordErrors = append(viewData.ConfirmPasswordErrors, e.GetMessage(ctr.ctx.TS.T(locale, "page.register.fields.confirm_password")))
				}
			}

		} else {
			if reqData.Password != reqData.ConfirmPassword {
				a := ctr.ctx.TS.T(locale, "page.register.fields.password")
				attributes := map[string]string{"attribute": a}
				viewData.PasswordErrors = append(viewData.PasswordErrors, "")
				viewData.ConfirmPasswordErrors = append(viewData.ConfirmPasswordErrors, ctr.ctx.TS.V(locale, "validation.confirmed", attributes))
			} else {
				valid = true
			}
		}

		if valid {
			userId, e := ctr.ctx.Auth.Register(reqData.Email, reqData.Password)
			if e != nil {
				if e.Code == "AuthService.Register.DuplicateUser" {
					status = http.StatusBadRequest
					viewData.Errors = append(viewData.Errors, ctr.ctx.TS.T(locale, "error.Auth.UserIsAlreadyRegistered"))
				} else {
					status = e.Status
					viewData.Errors = append(viewData.Errors, ctr.ctx.TS.T(locale, "error.500"))
				}
			} else {
				e = ctr.ctx.SS.Login(session, userId)
				if e != nil {
					status = e.Status
					viewData.Errors = append(viewData.Errors, ctr.ctx.TS.T(locale, "error.500"))
				} else {
					alert := dto.NewSuccessAlert(ctr.ctx.TS.T(locale, "alert.register.success"))
					//goland:noinspection GoUnhandledErrorResult
					ctr.ctx.AlS.AddFlash(session, []dto.Alert{alert})
					c.Redirect(http.StatusFound, "/login")
					return
				}
			}
		}

		viewData.EmailValue = reqData.Email
		viewData.PasswordValue = reqData.Password
		viewData.ConfirmPasswordValue = reqData.ConfirmPassword
	}

	d, err := view.NewRegister(c, ctr.ctx, &viewData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
