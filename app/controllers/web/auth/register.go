package auth

import (
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
	if user := session.Get(ctr.ctx.Cfg.Session.StoreUserKey); user != nil {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	locale := ctr.ctx.AS.Locale(c, nil)
	viewData := view.RegisterViewData{}
	reqData := RegisterRequest{}
	status := http.StatusOK
	if c.Request.Method == "POST" {
		// TODO: Сделать валидацию
		if err := c.ShouldBind(&reqData); err != nil {
			viewData.Errors = append(viewData.Errors, err.Error())
		} else {
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
				// TODO: Сделать Alert "Вы успешно зарегистрировались на сайте."
				e = ctr.ctx.SS.Login(session, userId)
				if e != nil {
					status = e.Status
					viewData.Errors = append(viewData.Errors, ctr.ctx.TS.T(locale, "error.500"))
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
