package services

import (
	"errors"
	"galaveg/app/dto"
	"github.com/gin-gonic/gin"
	"strings"
)

type LocaleService struct {
	localeCookieKey string
	defaultLocale   dto.Locale
	localesMap      map[string]dto.Locale
	localesCodes    []string
	locales         []dto.Locale
}

func NewLocaleService(defaultLocaleCode, localeCookieKey string, locales []dto.Locale) (*LocaleService, error) {
	localesMap := map[string]dto.Locale{}
	var localesCodes []string

	for _, l := range locales {
		localesCodes = append(localesCodes, l.Code)
		localesMap[l.Code] = l
	}
	defaultLocale, ok := localesMap[defaultLocaleCode]

	if !ok {
		return nil, errors.New("default locale not found")
	}

	return &LocaleService{
		localeCookieKey,
		defaultLocale,
		localesMap,
		localesCodes,
		locales,
	}, nil
}

func MustLocaleService(defaultLocaleCode, localeCookieKey string, locales []dto.Locale) *LocaleService {
	s, e := NewLocaleService(defaultLocaleCode, localeCookieKey, locales)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *LocaleService) GetLocale(code string) dto.Locale {
	l, _ := s.localesMap[code]
	return l
}

func (s *LocaleService) GetLocaleOrDefault(code string) dto.Locale {
	l, ok := s.localesMap[code]
	if ok {
		return l
	}
	return s.defaultLocale
}

func (s *LocaleService) GetDefault() dto.Locale {
	return s.defaultLocale
}

func (s *LocaleService) GetLocales() []dto.Locale {
	return s.locales
}

func (s *LocaleService) GetLocalesCodes() []string {
	return s.localesCodes
}

func (s *LocaleService) getExistsOrDefaultCode(code string) string {
	_, ok := s.localesMap[code]
	if ok {
		return code
	}
	return s.defaultLocale.Code
}

func (s *LocaleService) GetLocaleCodeFromGinContext(c *gin.Context, user *dto.User) string {
	var val string
	if c != nil {
		// manually selected by the user in the browser
		val, _ = c.Cookie(s.localeCookieKey)
		if val != "" {
			return s.getExistsOrDefaultCode(val)
		}
	}
	if user != nil {
		// the locale is selected from the database
		if user.Locale != "" {
			return s.getExistsOrDefaultCode(user.Locale)
		}
	}
	if c != nil {
		// Accept-Language: *
		// Accept-Language: fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5
		val = c.GetHeader("Accept-Language")
		if val != "" {
			if index := strings.IndexAny(val, "-,;"); index != -1 {
				return s.getExistsOrDefaultCode(val[:index])
			}
		}
	}
	return s.defaultLocale.Code
}

func (s *LocaleService) G(c *gin.Context, user *dto.User) string {
	return s.GetLocaleCodeFromGinContext(c, user)
}
