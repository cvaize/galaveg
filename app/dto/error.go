package dto

import "net/http"

type Error struct {
	Code    string `json:"code"`    // machine-readable error code
	Message string `json:"message"` // human-readable message
	Status  int    `json:"-"`       // HTTP status (not output in the body)
}

func (e *Error) Error() string {
	return e.Message
}

func NewBadRequest(code, message string) *Error {
	return &Error{Code: code, Message: message, Status: http.StatusBadRequest}
}

func NewNotFound(code, message string) *Error {
	return &Error{Code: code, Message: message, Status: http.StatusNotFound}
}

func NewInternal(code, message string) *Error {
	return &Error{Code: code, Message: message, Status: http.StatusInternalServerError}
}
