package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"github.com/gin-gonic/gin"
	"strings"
)

type AppService struct {
	config        config.AppConfig
	localeService *LocaleService
}

func NewAppService(config config.AppConfig, localeService *LocaleService) (*AppService, error) {
	return &AppService{config, localeService}, nil
}

func MustAppService(config config.AppConfig, localeService *LocaleService) *AppService {
	s, e := NewAppService(config, localeService)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *AppService) DarkMode(c *gin.Context) string {
	var val string
	if c != nil {
		val, _ = c.Cookie(s.config.DarkModeCookieKey)
	}
	if val != "auto" && val != "dark" && val != "light" {
		val = "auto"
	}

	return val
}

func (s *AppService) Locale(c *gin.Context, user *dto.User) string {
	var val string
	if c != nil {
		// manually selected by the user in the browser
		val, _ = c.Cookie(s.config.LocaleCookieKey)
		if val != "" {
			if s.localeService.Exists(val) {
				return val
			} else {
				return s.config.Locale
			}
		}
	}
	if user != nil && user.Locale != "" {
		// the locale is selected from the database
		if s.localeService.Exists(user.Locale) {
			return user.Locale
		} else {
			return s.config.Locale
		}
	}
	if c != nil {
		// Accept-Language: *
		// Accept-Language: fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5
		val = c.GetHeader("Accept-Language")
		if val != "" {
			if index := strings.IndexAny(val, "-,;"); index != -1 {
				if s.localeService.Exists(val[:index]) {
					return val[:index]
				} else {
					return s.config.Locale
				}
			}
		}
	}
	return s.config.Locale
}
