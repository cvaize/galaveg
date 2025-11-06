package services

import (
	"galaveg/app/services"
	"galaveg/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTranslatorService(t *testing.T) {
	c := tests.GetConfig()
	f := c.GetFolder("resources/lang/")

	_, err := services.NewTranslatorServiceFromFiles(f, c.App.Locale)
	assert.NoError(t, err)
}
