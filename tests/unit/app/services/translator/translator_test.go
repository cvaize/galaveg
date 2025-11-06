package translator

import (
	"fmt"
	services "galaveg/app/services/translator"
	"galaveg/tests"
	"testing"
)

func TestCreateTranslatorService(t *testing.T) {
	c := tests.GetConfig()
	f := c.GetFolder("resources/lang/**/*.json")

	s := services.New(f, c.App.Locale)

	fmt.Println(c.App.Locale)
	fmt.Println(s)
}
