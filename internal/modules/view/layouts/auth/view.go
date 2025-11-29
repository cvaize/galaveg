package auth

import (
	"galaveg/app/dto"
	"galaveg/app/view/components/btn"
	"galaveg/app/view/layouts/auth/form"
)

const TEMPLATE = "layouts/auth"

type View struct {
	Lang     string
	DarkMode string
	Title    string
	Heading  string
	Alerts   []dto.Alert
	Locale   dto.Locale
	Locales  []dto.Locale
	Form     form.View
	Back     *btn.View
}
