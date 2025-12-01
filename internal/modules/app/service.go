package app

import (
	"galaveg/internal/config"
	"galaveg/internal/modules/errors"
	"github.com/gin-gonic/gin"
	"net/url"
)

type Service = *ServiceImpl

type ServiceImpl struct {
	cfg *config.Config
	url *url.URL
}

func NewService(cfg *config.Config) (*ServiceImpl, *errors.Error) {
	u, err := url.Parse(cfg.Http.Url)
	if err != nil {
		return nil, errors.E500(err, "app.ServiceImpl.url.Parse", "")
	}
	return &ServiceImpl{cfg, u}, nil
}

func (s *ServiceImpl) DarkMode(c *gin.Context) string {
	var val string
	if c != nil {
		val, _ = c.Cookie(s.cfg.App.DarkModeCookieKey)
	}
	if val != "auto" && val != "dark" && val != "light" {
		val = "auto"
	}

	return val
}

func (s *ServiceImpl) Url() string {
	return s.url.String()
}

func (s *ServiceImpl) RefUrl() *url.URL {
	return s.url
}

func (s *ServiceImpl) CloneUrl() url.URL {
	return *s.url
}

func (s *ServiceImpl) LogoSrc() string {
	cloneUrl := s.CloneUrl()
	return cloneUrl.JoinPath("/svg/logo.svg").String()
}
