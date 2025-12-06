package errors

import (
	"galaveg/pkg/logger"
	"net/http"
)

// Translate keys

const Translate500 = "error.500"

func E400(cause error, code, message string) *Error {
	logger.Infof("(400) %s: %v", code, cause)
	return &Error{Code: code, Message: message, Status: http.StatusBadRequest}
}

func E401(cause error, code, message string) *Error {
	logger.Infof("(401) %s: %v", code, cause)
	return &Error{Code: code, Message: message, Status: http.StatusUnauthorized}
}

func E403(cause error, code, message string) *Error {
	logger.Infof("(403) %s: %v", code, cause)
	return &Error{Code: code, Message: message, Status: http.StatusForbidden}
}

func E404(cause error, code, message string) *Error {
	logger.Infof("(404) %s: %v", code, cause)
	return &Error{Code: code, Message: message, Status: http.StatusNotFound}
}

func E500(cause error, code, message string) *Error {
	logger.Errorf("(500) %s: %v", code, cause)
	return &Error{Code: code, Message: message, Status: http.StatusInternalServerError}
}
