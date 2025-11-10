package services

import "galaveg/app/dto"

type LocaleService struct {
	locale          string
	localeCookieKey string
	localesMap      map[string]dto.Locale
	localesCodes    []string
	locales         []dto.Locale
}

func NewLocaleService(locale, localeCookieKey string, locales []dto.Locale) *LocaleService {

	localesMap := map[string]dto.Locale{}
	var localesCodes []string

	for _, l := range locales {
		localesCodes = append(localesCodes, l.Code)
		localesMap[l.Code] = l
	}
	return &LocaleService{locale, localeCookieKey, localesMap, localesCodes, locales}
}
