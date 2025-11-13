package home

import (
	"galaveg/app/dto"
	"galaveg/app/services"
	"galaveg/app/view/components/breadcrumbs/item"
	"galaveg/app/view/components/sidebar"
	"galaveg/app/view/components/sidebar/brand"
	"galaveg/app/view/components/sidebar/menu_item"
	"github.com/gin-gonic/gin"
	"strings"
)

type View struct {
	Lang        string
	DarkMode    string
	Title       string
	Heading     string
	Breadcrumbs []item.View
	Sidebar     sidebar.View
	Alerts      []dto.Alert
}

func New(as *services.AppService, ls *services.LocaleService, ts *services.TranslatorService, c *gin.Context, user *dto.User) (View, error) {
	locale := ls.GetLocale(as.Locale(c, nil))
	locales := ls.GetLocales()
	path := c.FullPath()

	isProfileActive := strings.HasPrefix(path, "/profile")
	isUsersActive := strings.HasPrefix(path, "/users")
	isRolesActive := strings.HasPrefix(path, "/roles")
	menu := []menu_item.View{
		{
			Name: "home",
			Href: "/",
			Text: ts.T(locale.Code, "layout.sidebar.home"),
		},
		{
			Name:     "profile",
			Text:     user.Email,
			IsActive: isProfileActive,
			Dropdown: []menu_item.View{
				{
					Name:     "profile",
					Href:     "/profile",
					Text:     ts.T(locale.Code, "layout.sidebar.profile"),
					IsActive: isProfileActive,
				},
				{
					Name: "logout",
					Text: ts.T(locale.Code, "layout.sidebar.logout"),
				},
			},
		},
		{
			Name:     "users",
			Text:     ts.T(locale.Code, "layout.sidebar.users.index"),
			IsActive: isUsersActive && isRolesActive,
			Dropdown: []menu_item.View{
				{
					Name:     "users",
					Href:     "/users",
					Text:     ts.T(locale.Code, "layout.sidebar.users.index"),
					IsActive: isUsersActive,
				},
				{
					Name:     "roles",
					Href:     "/roles",
					Text:     ts.T(locale.Code, "layout.sidebar.users.roles"),
					IsActive: isRolesActive,
				},
			},
		},
		{
			Name:     "files",
			Href:     "/files",
			Text:     ts.T(locale.Code, "layout.sidebar.files"),
			IsActive: strings.HasPrefix(path, "/files"),
		},
	}

	localesMenuItem := menu_item.View{
		Name: "locales",
		Text: locale.FullName,
	}
	for _, v := range locales {
		if v.Code != locale.Code {
			localesMenuItem.Dropdown = append(localesMenuItem.Dropdown, menu_item.View{
				Text:  v.FullName,
				Value: v.Code,
			})
		}
	}
	menu = append(menu, localesMenuItem)
	return View{
		Lang:     locale.Code,
		DarkMode: as.DarkMode(c),
		Title:    ts.T(locale.Code, "page.home.title"),
		Heading:  ts.T(locale.Code, "page.home.header"),
		Breadcrumbs: []item.View{
			{
				Text: ts.T(locale.Code, "page.home.breadcrumbs.home"),
				Href: "/",
			},
		},
		Sidebar: sidebar.View{
			Brand: brand.View{
				Text:  ts.T(locale.Code, "layout.brand"),
				Href:  "/",
				Image: "/svg/logo.svg",
			},
			Menu: menu,
		},
		Alerts: as.Alerts(c),
	}, nil
}
