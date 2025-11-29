package form

import (
	"galaveg/internal/modules/view/components/btn"
	"galaveg/internal/modules/view/components/field"
)

type View struct {
	Action        string
	Method        string
	Fields        []field.View
	Submit        *btn.View
	ResetPassword *btn.View
	Register      *btn.View
	Login         *btn.View
	Errors        []string
	Text          string
}
