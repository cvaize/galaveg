package locales

import (
	"galaveg/internal/config"
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/users"
	"github.com/gin-gonic/gin"
	"slices"
	"strings"
)

type Service = *ServiceImpl

// TODO: Использовать cfg.Locales.Default вместо s.cfg.App.Locale и cfg.Locales.CookieKey вместо s.cfg.App.LocaleCookieKey

type ServiceImpl struct {
	cfg          *config.Config
	localesMap   map[string]Locale
	localesCodes []string
	locales      []Locale
}

func NewService(cfg *config.Config, locales []Locale) (*ServiceImpl, *errors.Error) {
	localesMap := map[string]Locale{}
	var localesCodes []string

	for _, l := range locales {
		localesCodes = append(localesCodes, l.Code)
		localesMap[l.Code] = l
	}

	return &ServiceImpl{
		cfg,
		localesMap,
		localesCodes,
		locales,
	}, nil
}

func (s *ServiceImpl) Exists(code string) bool {
	return slices.Contains(s.localesCodes, code)
}

func (s *ServiceImpl) GetLocale(code string) Locale {
	l, _ := s.localesMap[code]
	return l
}

func (s *ServiceImpl) GetLocales() []Locale {
	return s.locales
}

func (s *ServiceImpl) GetLocalesCodes() []string {
	return s.localesCodes
}

func (s *ServiceImpl) Locale(c *gin.Context, user *users.User) string {
	var val string
	if c != nil {
		// manually selected by the user in the browser
		val, _ = c.Cookie(s.cfg.App.LocaleCookieKey)
		if val != "" {
			if s.Exists(val) {
				return val
			} else {
				return s.cfg.App.Locale
			}
		}
	}
	if user != nil && user.Locale != "" {
		// the locale is selected from the database
		if s.Exists(user.Locale) {
			return user.Locale
		} else {
			return s.cfg.App.Locale
		}
	}
	if c != nil {
		// Accept-Language: *
		// Accept-Language: fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5
		val = c.GetHeader("Accept-Language")
		if val != "" {
			if index := strings.IndexAny(val, "-,;"); index != -1 {
				if s.Exists(val[:index]) {
					return val[:index]
				} else {
					return s.cfg.App.Locale
				}
			}
		}
	}
	return s.cfg.App.Locale
}
