package sidebar

import (
	"galaveg/app/dto"
	"galaveg/app/view/components/sidebar/brand"
	"galaveg/app/view/components/sidebar/menu_item"
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type View struct {
	Brand brand.View
	Menu  []menu_item.View
}

func New(c *gin.Context, user *dto.User) (*View, error) {
	locale := singleton.LS.GetLocale(singleton.AS.Locale(c, user))
	locales := singleton.LS.GetLocales()
	path := c.FullPath()

	isProfileActive := strings.HasPrefix(path, "/profile")
	isUsersActive := strings.HasPrefix(path, "/users")
	isRolesActive := strings.HasPrefix(path, "/roles")
	profileText := "Профиль"
	if user != nil {
		profileText = user.Email
	}
	menu := []menu_item.View{
		{
			Name: "home",
			Href: "/",
			Text: singleton.TS.T(locale.Code, "layout.sidebar.home"),
		},
		{
			Name:     "profile_dropdown",
			Text:     profileText,
			IsActive: isProfileActive,
			Dropdown: []menu_item.View{
				{
					Name:     "profile",
					Href:     "/profile",
					Text:     singleton.TS.T(locale.Code, "layout.sidebar.profile"),
					IsActive: isProfileActive,
				},
				{
					Name: "logout",
					Text: singleton.TS.T(locale.Code, "layout.sidebar.logout"),
				},
			},
			DropdownMaxHeight: "4rem",
		},
		{
			Name:     "users_dropdown",
			Text:     singleton.TS.T(locale.Code, "layout.sidebar.users.index"),
			IsActive: isUsersActive && isRolesActive,
			Dropdown: []menu_item.View{
				{
					Name:     "users",
					Href:     "/users",
					Text:     singleton.TS.T(locale.Code, "layout.sidebar.users.index"),
					IsActive: isUsersActive,
				},
				{
					Name:     "roles",
					Href:     "/roles",
					Text:     singleton.TS.T(locale.Code, "layout.sidebar.users.roles"),
					IsActive: isRolesActive,
				},
			},
			DropdownMaxHeight: "4rem",
		},
		{
			Name:     "files",
			Href:     "/files",
			Text:     singleton.TS.T(locale.Code, "layout.sidebar.files"),
			IsActive: strings.HasPrefix(path, "/files"),
		},
	}

	localesMenuItem := menu_item.View{
		Name:              "locales_dropdown",
		Text:              locale.FullName,
		DropdownMaxHeight: strconv.Itoa(len(locales)*2) + "rem",
	}

	for _, v := range locales {
		if v.Code != locale.Code {
			localesMenuItem.Dropdown = append(localesMenuItem.Dropdown, menu_item.View{
				Name:  "locale",
				Text:  v.FullName,
				Value: v.Code,
			})
		}
	}
	menu = append(menu, localesMenuItem)

	return &View{
		Brand: brand.View{
			Text:  singleton.TS.T(locale.Code, "layout.brand"),
			Href:  "/",
			Image: "",
		},
		Menu: menu,
	}, nil
}
