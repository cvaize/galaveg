package auth

import (
	"galaveg/app/view/layouts/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctr *Controller) ResetPasswordConfirm(c *gin.Context) {
	session := ctr.ctx.SS.Default(c)
	if ctr.ctx.SS.ExistsUserId(session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}
	d, err := auth.NewResetPasswordConfirm(c, ctr.ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, auth.TEMPLATE, d)
}
