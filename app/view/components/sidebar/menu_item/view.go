package menu_item

type View struct {
	Name     string
	Text     string
	Href     string
	Value    string
	IsActive bool
	Dropdown []View
}
