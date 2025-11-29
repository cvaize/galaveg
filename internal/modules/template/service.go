package template

import (
	"bytes"
	"galaveg/internal/modules/errors"
	htmlLib "html/template"
)

type Service struct {
	html *htmlLib.Template
}

func NewService(html *htmlLib.Template) (*Service, error) {
	return &Service{html}, nil
}

func (s *Service) Html(template string, data any) (string, *errors.Error) {
	var tpl bytes.Buffer
	if templateError := s.html.ExecuteTemplate(&tpl, template, data); templateError != nil {
		//goland:noinspection GoUnhandledErrorResult
		return "", errors.E500(templateError, "template.Service.Html.ExecuteTemplate", "")
	}

	return tpl.String(), nil
}
