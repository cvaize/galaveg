package home

import (
	"galaveg/app/dto"
	"galaveg/app/view/components/sidebar"
)

type View struct {
	Lang     string
	Title    string
	DarkMode string
	Csrf     string
	// "http://localhost:3000/"
	SiteUrl string
	// "/" or "/users/"
	Path    string
	Sidebar sidebar.View
	Alerts  []dto.Alert
	User    *dto.User
}
