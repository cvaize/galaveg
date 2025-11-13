package home

import (
	"galaveg/app/dto"
	"galaveg/app/view/components/sidebar"
)

type View struct {
	Lang        string
	DarkMode    string
	Csrf        string
	SiteUrl     string
	Path        string
	Title       string
	Heading     string
	Breadcrumbs string
	Sidebar     sidebar.View
	Alerts      []dto.Alert
	User        *dto.User
}
