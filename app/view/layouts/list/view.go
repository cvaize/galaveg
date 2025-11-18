package list

import (
	"galaveg/app/dto"
	"galaveg/app/view/components/breadcrumbs/item"
	"galaveg/app/view/components/sidebar"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
)

const TEMPLATE = "layouts/list"

type View struct {
	Lang     string
	DarkMode string
	//Csrf        string
	//SiteUrl     string
	//Path        string
	Title       string
	Heading     string
	Breadcrumbs []item.View
	Sidebar     *sidebar.View
	Alerts      []dto.Alert
}

func New(c *gin.Context, ctx *providers.Context, user *dto.User) (*View, error) {
	locale := ctx.LS.GetLocale(ctx.AS.Locale(c, user))

	sidebarObject, err := sidebar.New(c, ctx, user)

	if err != nil {
		return nil, err
	}

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.AS.DarkMode(c),
		Title:    ctx.TS.T(locale.Code, "page.home.title"),
		Heading:  ctx.TS.T(locale.Code, "page.home.header"),
		Breadcrumbs: []item.View{
			{
				Text: ctx.TS.T(locale.Code, "page.home.breadcrumbs.home"),
				Href: "/",
			},
			{
				Text: "Пользователи",
				Href: "/users",
			},
		},
		Sidebar: sidebarObject,
		Alerts:  ctx.AS.Alerts(c),
	}, nil
}
