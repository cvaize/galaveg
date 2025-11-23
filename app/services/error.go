package services

import (
	"galaveg/app/dto"
	"galaveg/utils/logger"
	"net/http"
)

type ErrorService struct {
}

func (s *ErrorService) E400(cause error, code, message string) *dto.Error {
	logger.Infof("(400) %s: %v", code, cause)
	return &dto.Error{Code: code, Message: message, Status: http.StatusBadRequest}
}

func (s *ErrorService) E401(cause error, code, message string) *dto.Error {
	logger.Infof("(401) %s: %v", code, cause)
	return &dto.Error{Code: code, Message: message, Status: http.StatusUnauthorized}
}

func (s *ErrorService) E403(cause error, code, message string) *dto.Error {
	logger.Infof("(403) %s: %v", code, cause)
	return &dto.Error{Code: code, Message: message, Status: http.StatusForbidden}
}

func (s *ErrorService) E404(cause error, code, message string) *dto.Error {
	logger.Infof("(404) %s: %v", code, cause)
	return &dto.Error{Code: code, Message: message, Status: http.StatusNotFound}
}

func (s *ErrorService) E500(cause error, code, message string) *dto.Error {
	logger.Errorf("(500) %s: %v", code, cause)
	return &dto.Error{Code: code, Message: message, Status: http.StatusInternalServerError}
}
