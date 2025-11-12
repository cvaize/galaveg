package web

import (
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	ctx, err := singleton.AS.NewWebDataCtx(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	d := singleton.AS.GetWebData(&ctx)
	l := ctx.Locale.Code
	d["Title"] = singleton.TS.T(l, "page.home.title")
	d["Heading"] = singleton.TS.T(l, "page.home.header")
	d["Breadcrumbs"] = []map[string]string{
		{
			"Label": singleton.TS.T(l, "page.home.breadcrumbs.home"),
			"Href":  "/",
		},
	}

	c.HTML(http.StatusOK, "layouts/home", d)
}
