package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"galaveg/app/dto"
	"galaveg/app/services"
)

func getLocales() []dto.Locale {
	return []dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
		{Code: "fr", ShortName: "fr", FullName: "Français"},
	}
}

func TestNewLocaleService(t *testing.T) {
	locales := getLocales()
	s, err := services.NewLocaleService(locales)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	// Проверяем корректную инициализацию
	assert.ElementsMatch(t, []string{"en", "ru", "fr"}, s.GetLocalesCodes())
	assert.Equal(t, locales, s.GetLocales())
}

func TestGetLocaleAndDefault(t *testing.T) {
	locales := getLocales()
	s, _ := services.NewLocaleService(locales)

	assert.Equal(t, "English", s.GetLocale("en").FullName)
	assert.Equal(t, dto.Locale{}, s.GetLocale("xx"))

	assert.Equal(t, "Русский", s.GetLocale("ru").FullName)
	assert.Equal(t, "English", s.GetLocale("en").FullName)
}
