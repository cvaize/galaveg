package app

import (
	"galaveg/config"
	"github.com/gin-gonic/gin"
	"net/url"
)

type Service struct {
	cfg *config.Config
	url *url.URL
}

func NewService(cfg *config.Config) (*Service, error) {
	u, err := url.Parse(cfg.Http.Url)
	if err != nil {
		return nil, err
	}
	return &Service{cfg, u}, nil
}

func (s *Service) DarkMode(c *gin.Context) string {
	var val string
	if c != nil {
		val, _ = c.Cookie(s.cfg.App.DarkModeCookieKey)
	}
	if val != "auto" && val != "dark" && val != "light" {
		val = "auto"
	}

	return val
}

func (s *Service) Url() string {
	return s.url.String()
}

func (s *Service) RefUrl() *url.URL {
	return s.url
}

func (s *Service) CloneUrl() url.URL {
	return *s.url
}

func (s *Service) LogoSrc() string {
	cloneUrl := s.CloneUrl()
	return cloneUrl.JoinPath("/svg/logo.svg").String()
}
