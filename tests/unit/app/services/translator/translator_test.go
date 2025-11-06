package translator

import (
	"fmt"
	"galaveg/config"
	"testing"
)

func TestCreateTranslatorService(t *testing.T) {
	c := config.New()
	f := c.GetFolder("resources/lang/**/*.json")

	fmt.Println(f)
	//fmt.Println(path.Path("/awda/awd"))
	//filenames, err := filepath.Glob("resources/lang/**/*.json")
	//filenames, err := filepath.Glob("./*")
	//fmt.Println(filenames)
	//fmt.Println(err)

	//if err != nil {
	//	return nil, err
	//}
	//if len(filenames) == 0 {
	//	return nil, fmt.Errorf("html/template: pattern matches no files: %#q", pattern)
	//}
	//repo := services.NewTranslatorService("")
	//
	//user := models.User{Name: "Alice", Email: "alice@example.com"}
	//created, err := repo.Create(user)
	//
	//if err != nil {
	//	t.Fatalf("expected no error, got %v", err)
	//}
	//if created.ID == 0 {
	//	t.Errorf("expected user ID to be set")
	//}
}
