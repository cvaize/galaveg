package web

import (
	"galaveg/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebController struct {
	AS *services.AppService
	TS *services.TranslatorService
}

func NewWebController(AS *services.AppService, TS *services.TranslatorService) (*WebController, error) {
	return &WebController{AS, TS}, nil
}

func MustNewWebController(AS *services.AppService, TS *services.TranslatorService) *WebController {
	ctr, err := NewWebController(AS, TS)
	if err != nil {
		panic(err)
	}
	return ctr
}

func (ctr *WebController) Home(c *gin.Context) {
	ctx, err := ctr.AS.NewWebDataCtx(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	d := ctr.AS.GetWebData(&ctx)
	l := ctx.Locale.Code
	d["Title"] = ctr.TS.T(l, "page.home.title")
	d["Heading"] = ctr.TS.T(l, "page.home.header")
	d["Breadcrumbs"] = []map[string]string{
		{
			"Label": ctr.TS.T(l, "page.home.breadcrumbs.home"),
			"Href":  "/",
		},
	}

	c.HTML(http.StatusOK, "layouts/home", d)
}
