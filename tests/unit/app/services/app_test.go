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

func GetAppService(t *testing.T) *services.AppService {
	cfg := tests.GetConfig()

	TS, e := services.NewTranslatorServiceFromFiles(cfg.GetFolder("resources/translates/"), cfg.App.Locale)
	assert.NoError(t, e)

	LS, e := services.NewLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
		{Code: "fr", ShortName: "fr", FullName: "Français"},
	})
	assert.NoError(t, e)

	RS, e := services.NewRoleService()
	assert.NoError(t, e)

	AS, e := services.NewAppService(cfg, LS, RS, TS)
	assert.NoError(t, e)

	return AS
}

func TestNewAppService(t *testing.T) {
	s := GetAppService(t)
	assert.NotNil(t, s)
}

func TestLang(t *testing.T) {
	s := GetAppService(t)

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

	s := GetAppService(t)

	assert.Equal(t, "fr", s.Locale(c, nil))
}

func TestLang_UserLocale(t *testing.T) {
	s := GetAppService(t)

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

	s := GetAppService(t)

	assert.Equal(t, "fr", s.Locale(c, nil))
}

func TestLang_Default(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	s := GetAppService(t)

	assert.Equal(t, "en", s.Locale(c, nil))
}
