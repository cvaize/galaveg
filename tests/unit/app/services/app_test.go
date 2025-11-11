package services

import (
	"galaveg/app/dto"
	"galaveg/tests"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"galaveg/app/services"
)

func GetAppService() (*services.AppService, error) {
	cf := tests.GetConfig()
	ls, e := services.NewLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
		{Code: "fr", ShortName: "fr", FullName: "Français"},
	})
	if e != nil {
		return nil, e
	}
	return services.NewAppService(cf.App, ls)
}

func TestNewAppService(t *testing.T) {
	s, err := GetAppService()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestLang(t *testing.T) {
	s, _ := GetAppService()

	assert.Equal(t, "ru", s.Locale(nil, &dto.User{Locale: "ru"}))
	assert.Equal(t, "en", s.Locale(nil, &dto.User{Locale: "xx"}))
}

func TestLang_Cookie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cf := tests.GetConfig()

	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Header.Set("Cookie", cf.App.LocaleCookieKey+"=fr")

	s, _ := GetAppService()

	assert.Equal(t, "fr", s.Locale(c, nil))
}

func TestLang_UserLocale(t *testing.T) {
	s, _ := GetAppService()

	assert.Equal(t, "ru", s.Locale(nil, &dto.User{Locale: "ru"}))
	assert.Equal(t, "en", s.Locale(nil, &dto.User{Locale: "xx"}))
}

func TestLang_Header(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Header.Set("Accept-Language", "fr-CH,fr;q=0.9,en;q=0.8")

	s, _ := GetAppService()

	assert.Equal(t, "fr", s.Locale(c, nil))
}

func TestLang_Default(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	s, _ := GetAppService()

	assert.Equal(t, "en", s.Locale(c, nil))
}
