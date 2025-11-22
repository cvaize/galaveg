package auth

import (
	"galaveg/app/view/layouts/auth"
	"galaveg/bootstrap/providers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	ctx *providers.Context
}

func NewController(ctx *providers.Context) Controller {
	return Controller{ctx}
}

func (ctr *Controller) Login(c *gin.Context) {
	session := sessions.Default(c)
	if user := session.Get(ctr.ctx.Cfg.Session.StoreUserKey); user != nil {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	//email := c.PostForm("email")
	//password := c.PostForm("password")

	//userId := "123"
	//
	//// Save the username in the session
	//session.Set(ctr.ctx.Cfg.Session.StoreUserKey, userId) // In real world usage you'd set this to the users ID
	//if err := session.Save(); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
	//	return
	//}

	d, err := auth.NewLogin(c, ctr.ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, auth.TEMPLATE, d)
}
