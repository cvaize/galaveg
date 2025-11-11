package services

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"galaveg/app/dto"
	"galaveg/app/services"
)

const LOCALE_COOKIE_KEY = "locale"

func getLocales() []dto.Locale {
	return []dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
		{Code: "fr", ShortName: "fr", FullName: "Français"},
	}
}

func TestNewLocaleService(t *testing.T) {
	locales := getLocales()
	s, err := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	// Проверяем корректную инициализацию
	assert.Equal(t, "en", s.GetDefault().Code)
	assert.ElementsMatch(t, []string{"en", "ru", "fr"}, s.GetLocalesCodes())
	assert.Equal(t, locales, s.GetLocales())
}

func TestNewLocaleService_DefaultNotFound(t *testing.T) {
	locales := getLocales()
	s, err := services.NewLocaleService("de", LOCALE_COOKIE_KEY, locales)
	assert.Nil(t, s)
	assert.Error(t, err)
	assert.Equal(t, errors.New("default locale not found").Error(), err.Error())
}

func TestMustLocaleService(t *testing.T) {
	locales := getLocales()
	assert.NotPanics(t, func() {
		_ = services.MustLocaleService("en", LOCALE_COOKIE_KEY, locales)
	})

	assert.Panics(t, func() {
		_ = services.MustLocaleService("de", LOCALE_COOKIE_KEY, locales)
	})
}

func TestGetLocaleAndDefault(t *testing.T) {
	locales := getLocales()
	s, _ := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)

	assert.Equal(t, "English", s.GetLocale("en").FullName)
	assert.Equal(t, dto.Locale{}, s.GetLocale("xx"))

	assert.Equal(t, "Русский", s.GetLocaleOrDefault("ru").FullName)
	assert.Equal(t, "English", s.GetLocaleOrDefault("xx").FullName)
}

func TestGetExistsOrDefaultCode(t *testing.T) {
	locales := getLocales()
	s, _ := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)

	assert.Equal(t, "ru", s.GetLocaleCodeFromGinContext(nil, &dto.User{Locale: "ru"}))
	assert.Equal(t, "en", s.GetLocaleCodeFromGinContext(nil, &dto.User{Locale: "xx"}))
}

func TestGetLocaleCodeFromGinContext_Cookie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Header.Set("Cookie", LOCALE_COOKIE_KEY+"=fr")

	locales := getLocales()
	s, _ := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)

	assert.Equal(t, "fr", s.GetLocaleCodeFromGinContext(c, nil))
}

func TestGetLocaleCodeFromGinContext_UserLocale(t *testing.T) {
	locales := getLocales()
	s, _ := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)

	assert.Equal(t, "ru", s.GetLocaleCodeFromGinContext(nil, &dto.User{Locale: "ru"}))
	assert.Equal(t, "en", s.GetLocaleCodeFromGinContext(nil, &dto.User{Locale: "xx"}))
}

func TestGetLocaleCodeFromGinContext_Header(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Header.Set("Accept-Language", "fr-CH,fr;q=0.9,en;q=0.8")

	locales := getLocales()
	s, _ := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)

	assert.Equal(t, "fr", s.GetLocaleCodeFromGinContext(c, nil))
}

func TestGetLocaleCodeFromGinContext_Default(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	locales := getLocales()
	s, _ := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)

	assert.Equal(t, "en", s.GetLocaleCodeFromGinContext(c, nil))
}

func TestG(t *testing.T) {
	locales := getLocales()
	s, _ := services.NewLocaleService("en", LOCALE_COOKIE_KEY, locales)
	assert.Equal(t, "en", s.G(nil, nil))
}
