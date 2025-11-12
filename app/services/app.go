package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"github.com/gin-gonic/gin"
	"strings"
)

type AppService struct {
	config config.AppConfig
	ls     *LocaleService
	rs     *RoleService
	ts     *TranslatorService
}

func NewAppService(config config.AppConfig, ls *LocaleService, rs *RoleService, ts *TranslatorService) (*AppService, error) {
	return &AppService{config, ls, rs, ts}, nil
}

func MustAppService(config config.AppConfig, ls *LocaleService, rs *RoleService, ts *TranslatorService) *AppService {
	s, e := NewAppService(config, ls, rs, ts)
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
			if s.ls.Exists(val) {
				return val
			} else {
				return s.config.Locale
			}
		}
	}
	if user != nil && user.Locale != "" {
		// the locale is selected from the database
		if s.ls.Exists(user.Locale) {
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
				if s.ls.Exists(val[:index]) {
					return val[:index]
				} else {
					return s.config.Locale
				}
			}
		}
	}
	return s.config.Locale
}

func (s *AppService) Csrf(c *gin.Context) string {
	return ""
}

func (s *AppService) Alerts(c *gin.Context) []dto.Alert {
	return []dto.Alert{}
}

type WebDataCtx struct {
	User     *dto.User
	Locale   dto.Locale
	Locales  []dto.Locale
	Alerts   []dto.Alert
	Path     string
	Title    string
	Heading  string
	DarkMode string
	Csrf     string
}

func (s *AppService) NewWebDataCtx(c *gin.Context) (WebDataCtx, error) {
	locale := s.ls.GetLocale(s.Locale(c, nil))
	locales := s.ls.GetLocales()

	return WebDataCtx{
		User:     nil,
		Locale:   locale,
		Locales:  locales,
		Alerts:   s.Alerts(c),
		Path:     c.FullPath(),
		Title:    s.ts.T(locale.Code, "app.name"),
		DarkMode: s.DarkMode(c),
		Csrf:     s.Csrf(c),
	}, nil
}

func (s *AppService) GetWebData(ctx *WebDataCtx) gin.H {
	return gin.H{
		"Lang":     ctx.Locale.Code,
		"Title":    ctx.Title,
		"DarkMode": ctx.DarkMode,
	}
}
