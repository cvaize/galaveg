package list

import (
	"galaveg/internal/modules/alerts"
	"galaveg/internal/modules/app"
	localesModule "galaveg/internal/modules/locales"
	"galaveg/internal/modules/translator"
	"galaveg/internal/modules/users"
	"galaveg/internal/modules/view/components/breadcrumbs/item"
	"galaveg/internal/modules/view/components/sidebar"
	"github.com/gin-contrib/sessions"
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
	Alerts      []alerts.Alert
}

func New(c *gin.Context, as app.Service, ls *localesModule.ServiceImpl, ts translator.Service, s sessions.Session, user *users.User) (*View, error) {
	locale := ls.GetLocale(ls.Locale(c, user))

	sidebarObject, err := sidebar.New(c, ls, ts, user)

	if err != nil {
		return nil, err
	}

	return &View{
		Lang:     locale.Code,
		DarkMode: as.DarkMode(c),
		Title:    ts.T(locale.Code, "page.home.title"),
		Heading:  ts.T(locale.Code, "page.home.header"),
		Breadcrumbs: []item.View{
			{
				Text: ts.T(locale.Code, "page.home.breadcrumbs.home"),
				Href: "/",
			},
			{
				Text: "Пользователи",
				Href: "/users",
			},
		},
		Sidebar: sidebarObject,
		Alerts:  alerts.Flashes(s),
	}, nil
}
