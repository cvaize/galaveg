package services

import (
	"galaveg/app/dto"
	"slices"
)

type LocaleService struct {
	localesMap   map[string]dto.Locale
	localesCodes []string
	locales      []dto.Locale
}

func NewLocaleService(locales []dto.Locale) (*LocaleService, error) {
	localesMap := map[string]dto.Locale{}
	var localesCodes []string

	for _, l := range locales {
		localesCodes = append(localesCodes, l.Code)
		localesMap[l.Code] = l
	}

	return &LocaleService{
		localesMap,
		localesCodes,
		locales,
	}, nil
}

func MustLocaleService(locales []dto.Locale) *LocaleService {
	s, e := NewLocaleService(locales)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *LocaleService) Exists(code string) bool {
	return slices.Contains(s.localesCodes, code)
}

func (s *LocaleService) GetLocale(code string) dto.Locale {
	l, _ := s.localesMap[code]
	return l
}

func (s *LocaleService) GetLocales() []dto.Locale {
	return s.locales
}

func (s *LocaleService) GetLocalesCodes() []string {
	return s.localesCodes
}
