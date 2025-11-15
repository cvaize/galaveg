package form

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
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
