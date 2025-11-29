package auth

import (
	"galaveg/internal/modules/alerts"
	"galaveg/internal/modules/locales"
	"galaveg/internal/modules/view/components/btn"
	"galaveg/internal/modules/view/layouts/auth/form"
)

const TEMPLATE = "layouts/auth"

type View struct {
	Lang     string
	DarkMode string
	Title    string
	Heading  string
	Alerts   []alerts.Alert
	Locale   locales.Locale
	Locales  []locales.Locale
	Form     form.View
	Back     *btn.View
}
