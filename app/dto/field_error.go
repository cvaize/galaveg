package dto

import "strings"

type FieldError struct {
	Name    string
	Message string
}

func (e *FieldError) GetMessage(name string) string {
	return strings.Replace(e.Message, e.Name, name, 1)
}
