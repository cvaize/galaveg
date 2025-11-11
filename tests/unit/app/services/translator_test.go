package services

import (
	"galaveg/app/services"
	"galaveg/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTranslates() map[string]map[string]string {
	return map[string]map[string]string{
		"en": {
			"hello":        "Hello",
			"world":        "World",
			"welcome":      "Welcome",
			"fallback":     "Fallback text",
			"only_en":      "Only in English",
			"welcome_vars": "Welcome :name to :city",
			"message_vars": "Hello :name, you have :count messages",
			"messages":     "message|messages",
			"apples":       "apple|apples",
		},
		"ru": {
			"hello":    "Привет",
			"world":    "Мир",
			"welcome":  "Добро пожаловать",
			"messages": "сообщение|сообщения|сообщений",
			"apples":   "яблоко|яблока|яблок",
		},
	}
}

func TestNewTranslatorService(t *testing.T) {
	translates := getTranslates()
	service := services.NewTranslatorService("en", translates)

	assert.NotNil(t, service)
	assert.Equal(t, service.GetLocale(), "en")
	assert.Equal(t, len(service.GetTranslates()), 2)
}

func TestNewTranslatorServiceFromFiles(t *testing.T) {
	c := tests.GetConfig()
	f := c.GetFolder("resources/translates/")

	_, err := services.NewTranslatorServiceFromFiles(f, c.App.Locale)
	assert.NoError(t, err)
}

func TestTranslatorService_Get(t *testing.T) {
	translates := getTranslates()
	service := services.NewTranslatorService("en", translates)

	assert.Equal(t, service.Get("en", "hello"), "Hello")
	assert.Equal(t, service.Get("ru", "hello"), "Привет")
	assert.Equal(t, service.Get("en", "nonexistent"), "")
	assert.Equal(t, service.Get("fr", "hello"), "")
}

func TestTranslatorService_Is(t *testing.T) {
	translates := getTranslates()
	service := services.NewTranslatorService("en", translates)

	assert.Equal(t, service.Is("en", "hello"), true)
	assert.Equal(t, service.Is("ru", "hello"), true)
	assert.Equal(t, service.Is("en", "nonexistent"), false)
	assert.Equal(t, service.Is("fr", "hello"), false)
}

func TestTranslatorService_Translate(t *testing.T) {
	translates := getTranslates()
	service := services.NewTranslatorService("en", translates)

	assert.Equal(t, service.T("en", "hello"), "Hello")
	assert.Equal(t, service.T("ru", "hello"), "Привет")
	assert.Equal(t, service.T("fr", "welcome"), "Welcome")
	assert.Equal(t, service.T("en", "nonexistent"), "nonexistent")
	assert.Equal(t, service.T("fr", "nonexistent"), "nonexistent")
}

func TestTranslatorService_Contains(t *testing.T) {
	translates := getTranslates()
	service := services.NewTranslatorService("en", translates)

	assert.Equal(t, service.Contains("ru", "hello"), true)
	assert.Equal(t, service.Contains("ru", "only_en"), true)
	assert.Equal(t, service.Contains("en", "nonexistent"), false)
	assert.Equal(t, service.Contains("fr", "hello"), true)
	assert.Equal(t, service.Contains("fr", "nonexistent"), false)
}

func TestTranslatorService_Variables(t *testing.T) {
	translates := getTranslates()
	service := services.NewTranslatorService("en", translates)

	assert.Equal(t, service.V("en", "welcome_vars", map[string]string{"name": "John"}), "Welcome John to :city")
	assert.Equal(t, service.V("en", "welcome_vars", map[string]string{"name": "John", "city": "London"}), "Welcome John to London")
	assert.Equal(t, service.V("en", "welcome_vars", nil), "Welcome :name to :city")
	assert.Equal(t, service.V("en", "welcome_vars", map[string]string{}), "Welcome :name to :city")
}

func TestTranslatorService_Choices(t *testing.T) {
	translates := getTranslates()
	service := services.NewTranslatorService("en", translates)

	assert.Equal(t, service.C("en", "messages", 1, nil), "message")
	assert.Equal(t, service.C("en", "messages", 2, nil), "messages")
	assert.Equal(t, service.C("en", "messages", 0, nil), "messages")
	assert.Equal(t, service.C("en", "messages", -1, nil), "message")

	assert.Equal(t, service.C("ru", "messages", 1, nil), "сообщение")
	assert.Equal(t, service.C("ru", "messages", 21, nil), "сообщение")
	assert.Equal(t, service.C("ru", "messages", 2, nil), "сообщения")
	assert.Equal(t, service.C("ru", "messages", 3, nil), "сообщения")
	assert.Equal(t, service.C("ru", "messages", 4, nil), "сообщения")
	assert.Equal(t, service.C("ru", "messages", 5, nil), "сообщений")
	assert.Equal(t, service.C("ru", "messages", 11, nil), "сообщений")
	assert.Equal(t, service.C("ru", "messages", 25, nil), "сообщений")
	assert.Equal(t, service.C("ru", "messages", 0, nil), "сообщений")
	assert.Equal(t, service.C("ru", "messages", 9999, nil), "сообщений")
	assert.Equal(t, service.C("ru", "messages", 78433, nil), "сообщения")
}
