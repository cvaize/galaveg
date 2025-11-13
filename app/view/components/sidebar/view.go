package sidebar

import (
	"galaveg/app/view/components/sidebar/brand"
	"galaveg/app/view/components/sidebar/menu_item"
)

type View struct {
	Brand brand.View
	Menu  []menu_item.View
}
