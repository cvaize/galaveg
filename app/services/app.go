package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

type AppService struct {
	cfg *config.Config
	ls  *LocaleService
	rs  *RoleService
	ts  *TranslatorService
	url *url.URL
}

func NewAppService(c *config.Config, ls *LocaleService, rs *RoleService, ts *TranslatorService) (*AppService, error) {
	u, err := url.Parse(c.App.Url)
	if err != nil {
		panic(err)
	}
	return &AppService{c, ls, rs, ts, u}, nil
}

func (s *AppService) DarkMode(c *gin.Context) string {
	var val string
	if c != nil {
		val, _ = c.Cookie(s.cfg.App.DarkModeCookieKey)
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
		val, _ = c.Cookie(s.cfg.App.LocaleCookieKey)
		if val != "" {
			if s.ls.Exists(val) {
				return val
			} else {
				return s.cfg.App.Locale
			}
		}
	}
	if user != nil && user.Locale != "" {
		// the locale is selected from the database
		if s.ls.Exists(user.Locale) {
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
				if s.ls.Exists(val[:index]) {
					return val[:index]
				} else {
					return s.cfg.App.Locale
				}
			}
		}
	}
	return s.cfg.App.Locale
}

func (s *AppService) Csrf(c *gin.Context) string {
	return ""
}

func (s *AppService) Alerts(c *gin.Context) []dto.Alert {
	return []dto.Alert{}
}

func (s *AppService) Url() string {
	return s.url.String()
}

func (s *AppService) RefUrl() *url.URL {
	return s.url
}

func (s *AppService) CloneUrl() url.URL {
	return *s.url
}

func (s *AppService) LogoSrc() string {
	cloneUrl := s.CloneUrl()
	return cloneUrl.JoinPath("/svg/logo.svg").String()
}
