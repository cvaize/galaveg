package auth

import (
	"fmt"
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

	viewData := view.RegisterViewData{}
	reqData := RegisterRequest{}
	if c.Request.Method == "POST" {
		// TODO: Сделать валидацию
		if err := c.ShouldBind(&reqData); err != nil {
			viewData.Errors = []string{err.Error()}
			fmt.Println(err)
		}

		viewData.EmailValue = reqData.Email
		viewData.PasswordValue = reqData.Password
		viewData.ConfirmPasswordValue = reqData.ConfirmPassword

		fmt.Println(reqData)
	}

	d, err := view.NewRegister(c, ctr.ctx, &viewData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, view.TEMPLATE, d)
}
