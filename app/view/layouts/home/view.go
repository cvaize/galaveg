package home

import (
	"galaveg/app/dto"
	"galaveg/app/view/components/breadcrumbs/item"
	"galaveg/app/view/components/sidebar"
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
)

type View struct {
	Lang        string
	DarkMode    string
	Title       string
	Heading     string
	Breadcrumbs []item.View
	Sidebar     *sidebar.View
	Alerts      []dto.Alert
}

func New(c *gin.Context, user *dto.User) (*View, error) {
	locale := singleton.LS.GetLocale(singleton.AS.Locale(c, nil))

	sidebarObject, err := sidebar.New(c, user)

	if err != nil {
		return nil, err
	}

	return &View{
		Lang:     locale.Code,
		DarkMode: singleton.AS.DarkMode(c),
		Title:    singleton.TS.T(locale.Code, "page.home.title"),
		Heading:  singleton.TS.T(locale.Code, "page.home.header"),
		Breadcrumbs: []item.View{
			{
				Text: singleton.TS.T(locale.Code, "page.home.breadcrumbs.home"),
				Href: "/",
			},
		},
		Sidebar: sidebarObject,
		Alerts:  singleton.AS.Alerts(c),
	}, nil
}
