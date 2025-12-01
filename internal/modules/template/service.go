package template

import (
	"bytes"
	"galaveg/internal/modules/errors"
	htmlLib "html/template"
)

type Service = *ServiceImpl

type ServiceImpl struct {
	html *htmlLib.Template
}

func NewService(html *htmlLib.Template) (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{html}, nil
}

func (s *ServiceImpl) Html(template string, data any) (string, *errors.Error) {
	var tpl bytes.Buffer
	if templateError := s.html.ExecuteTemplate(&tpl, template, data); templateError != nil {
		//goland:noinspection GoUnhandledErrorResult
		return "", errors.E500(templateError, "template.ServiceImpl.Html.ExecuteTemplate", "")
	}

	return tpl.String(), nil
}
