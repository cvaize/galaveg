package errors

import "strings"

type Error struct {
	Code    string `json:"code"`    // machine-readable error code
	Message string `json:"message"` // human-readable message
	Status  int    `json:"-"`       // HTTP status (not output in the body)
	//Cause   error  `json:"-"`       // original error (for logging)
}

func (e *Error) Error() string {
	return e.Message
}

type FieldError struct {
	Name    string
	Message string
}

func (e *FieldError) GetMessage(name string) string {
	return strings.Replace(e.Message, e.Name, name, 1)
}
