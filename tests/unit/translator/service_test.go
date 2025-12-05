package translator_test

import (
	"errors"
	"galaveg/pkg/logger"
	"os"
	"path/filepath"
	"testing"

	tr "galaveg/internal/modules/translator"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	translates := map[string]map[string]string{
		"en": {"hello": "Hello"},
		"ru": {"hello": "Привет"},
	}

	s, err := tr.NewService("en", translates)

	// The constructor must not return error
	assert.Nil(t, err)

	// Service must be initialized correctly
	assert.NotNil(t, s)
	assert.Equal(t, "en", s.GetLocale())
	assert.Equal(t, "Hello", s.T("en", "hello"))
	assert.Equal(t, "Привет", s.T("ru", "hello"))
}

func TestT(t *testing.T) {
	s, _ := tr.NewService("en", map[string]map[string]string{
		"en": {"greet": "Hello"},
	})

	// Must return translation
	assert.Equal(t, "Hello", s.T("en", "greet"))

	// Must return key if not found
	assert.Equal(t, "unknown.key", s.T("en", "unknown.key"))

	// Must return value for default locale
	assert.Equal(t, "Hello", s.T("xx", "greet"))
}

func TestNewServiceFromFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create translation files
	_ = os.WriteFile(filepath.Join(tmpDir, "en.json"), []byte(`{"hello": "Hello"}`), 0644)
	_ = os.WriteFile(filepath.Join(tmpDir, "ru.json"), []byte(`{"hello": "Привет"}`), 0644)

	s, err := tr.NewServiceFromFiles(tmpDir, "en")

	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "Hello", s.T("en", "hello"))
	assert.Equal(t, "Привет", s.T("ru", "hello"))
}

func TestNewServiceFromFiles_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()

	logger.SetLogLevel("panic")
	// Invalid JSON
	_ = os.WriteFile(filepath.Join(tmpDir, "en.json"), []byte(`{invalid json}`), 0644)

	s, err := tr.NewServiceFromFiles(tmpDir, "en")

	assert.Nil(t, s)
	assert.NotNil(t, err)
}

func TestNewServiceFromFiles_NoDirectory(t *testing.T) {
	logger.SetLogLevel("panic")
	s, err := tr.NewServiceFromFiles("path/that/does/not/exist", "en")

	assert.Nil(t, s)
	assert.NotNil(t, err)
}

func TestTVE_ValidationErrors(t *testing.T) {
	validate := validator.New()

	type TestData struct {
		Email string `validate:"required,email"`
	}

	// Prepare translator
	translates := map[string]map[string]string{
		"en": {
			"Email": "Email", // used for error messages
		},
	}

	s, _ := tr.NewService("en", translates)

	// Validate struct
	data := TestData{}
	err := validate.Struct(data)

	// Translate validation errors
	errs := s.TVE("en", err)

	assert.NotEmpty(t, errs, "Validation must produce translated errors")
	assert.Equal(t, "Email", errs[0].Name) // field name
	assert.NotEmpty(t, errs[0].Message)    // translated message text
}

func TestTVE_NotValidationError(t *testing.T) {
	s, _ := tr.NewService("en", map[string]map[string]string{})

	errs := s.TVE("en", errors.New("some error"))

	assert.Len(t, errs, 0, "Non-validation error must produce no translated errors")
}
