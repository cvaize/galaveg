package services

import (
	"bytes"
	"galaveg/config"
	"galaveg/utils/logger"
	htmlLib "html/template"
)

type TemplateService struct {
	cfg  *config.Config
	html *htmlLib.Template
}

func NewTemplateService(cfg *config.Config, html *htmlLib.Template) (*TemplateService, error) {
	return &TemplateService{cfg, html}, nil
}

func MustTemplateService(cfg *config.Config, html *htmlLib.Template) *TemplateService {
	s, e := NewTemplateService(cfg, html)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *TemplateService) Html(template string, data any) (string, error) {
	var tpl bytes.Buffer
	if templateError := s.html.ExecuteTemplate(&tpl, template, data); templateError != nil {
		logger.Errorf("(500) TemplateService.Html.ExecuteTemplate: %v", templateError)
		return "", templateError
	}

	return tpl.String(), nil
}
