package translator

import (
	"encoding/json"
	"errors"
	"fmt"
	errorsModule "galaveg/internal/modules/errors"
	"galaveg/pkg/flatten"
	"github.com/gin-gonic/gin/binding"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	ruTranslations "github.com/go-playground/validator/v10/translations/ru"
)

// Service is the interface for the translation service
type Service = *ServiceImpl

// ServiceImpl implements the translation service for internationalization (i18n)
// and validation error translation.
type ServiceImpl struct {
	locale     string                       // Current application locale (e.g., "en", "ru")
	translates map[string]map[string]string // Nested map: locale -> translation key -> translated value
	u          *ut.UniversalTranslator      // Universal translator for validation errors
}

// NewService creates a new translation service with the given locale and translation map.
func NewService(locale string, translates map[string]map[string]string) (*ServiceImpl, *errorsModule.Error) {
	// Initialize locale-specific translators
	enLocale := en.New()
	ruLocale := ru.New()

	// Create a universal translator for embedded go-playground/validator/v10 validation
	u := ut.New(enLocale, enLocale, ruLocale)

	// Register validation error translations for supported locales
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		transEN, _ := u.GetTranslator("en")
		_ = enTranslations.RegisterDefaultTranslations(v, transEN)

		transRU, _ := u.GetTranslator("ru")
		_ = ruTranslations.RegisterDefaultTranslations(v, transRU)
	}
	return &ServiceImpl{locale, translates, u}, nil
}

// NewServiceFromFiles creates a new translation service by loading translation files
// from the specified directory. Files should be in JSON format with locale as part of the filename.
func NewServiceFromFiles(dir, locale string) (*ServiceImpl, *errorsModule.Error) {
	translates := map[string]map[string]string{}

	// Clean and validate the directory path
	dir = filepath.Clean(dir)
	stat, err := os.Stat(dir)
	if err != nil || !stat.IsDir() {
		return nil, errorsModule.E500(err, "translator.ServiceImpl.NewServiceFromFiles",
			fmt.Sprintf("translates folder not found: %s", dir))
	}

	// Walk through the translation directory to load all JSON files
	var fullKey, locale1, prefixKey, key, value string
	var dotIndex int
	var valueAny interface{}
	err = filepath.WalkDir(dir, func(fullPath string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		// Process only JSON files
		if strings.HasSuffix(strings.ToLower(d.Name()), ".json") {
			// Convert file path to translation key format (e.g., "en.users.page" from "en/users/page.json")
			fullKey = strings.ReplaceAll(fullPath, dir, "")
			fullKey = strings.ReplaceAll(fullKey, string(filepath.Separator), ".")
			fullKey = strings.TrimLeft(fullKey, ".")
			fullKey = strings.ReplaceAll(fullKey, ".json", "")

			// Read and parse the JSON file
			content, err := os.ReadFile(fullPath)
			if err != nil {
				return err
			}

			var parsed map[string]interface{}
			if err := json.Unmarshal(content, &parsed); err != nil {
				return err
			}

			// Flatten the nested JSON structure to dot notation
			flat, err := flatten.Flatten(parsed, "", flatten.DotStyle)
			if err != nil {
				return err
			}

			if len(flat) != 0 {
				// Extract locale from the file path (first part of the key)
				dotIndex = strings.Index(fullKey, ".")
				locale1 = ""
				prefixKey = ""

				if dotIndex == -1 {
					locale1 = fullKey
				} else {
					locale1 = fullKey[:dotIndex]
					prefixKey = fullKey[dotIndex+1:]
				}

				locale1 = strings.TrimSpace(locale1)
				prefixKey = strings.TrimSpace(prefixKey)

				if locale1 == "" {
					return fmt.Errorf("the language is not defined by the path: %s", fullPath)
				}

				// Initialize locale map if it doesn't exist
				if _, ok := translates[locale1]; !ok {
					translates[locale1] = map[string]string{}
				}

				// Add all translations from this file to the locale map
				for key, valueAny = range flat {
					key = fmt.Sprintf("%s.%s", prefixKey, key)
					key = strings.TrimLeft(key, ".")
					key = strings.TrimSpace(key)
					value = strings.TrimSpace(fmt.Sprintf("%s", valueAny))
					translates[locale1][key] = value
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, errorsModule.E500(err, "translator.ServiceImpl.NewServiceFromFiles", "")
	}

	// Process translation variables (interpolation of {{variable}} references)
	var variables, sp1, sp2 []string
	var varSt, v string
	var ok bool
	for locale1, _ = range translates {
		for key, value = range translates[locale1] {
			variables = []string{}

			// Extract all variable references from the translation value
			sp1 = strings.Split(value, "{{")
			for _, varSt = range sp1 {
				sp2 = strings.Split(varSt, "}}")

				if len(sp2) > 1 {
					variables = append(variables, sp2[0])
				}
			}

			// Replace variable references with their actual values
			if len(variables) > 0 {
				for _, varSt = range variables {
					v, ok = translates[locale1][strings.TrimSpace(varSt)]
					if !ok {
						v, ok = translates[locale1][strings.TrimSpace(varSt)]
					}
					if ok {
						translates[locale1][key] = strings.ReplaceAll(translates[locale1][key], "{{"+varSt+"}}", v)
					}
				}
			}
		}
	}

	// Create the service with the loaded translations
	return NewService(locale, translates)
}

// GetLocale returns the current application locale
func (s *ServiceImpl) GetLocale() string {
	return s.locale
}

// GetTranslates returns the complete translation map
func (s *ServiceImpl) GetTranslates() map[string]map[string]string {
	return s.translates
}

// Get retrieves a translation for a specific locale and key
func (s *ServiceImpl) Get(locale, key string) string {
	v, _ := s.translates[locale][key]
	return v
}

// Is checks if a translation exists for the given locale and key
func (s *ServiceImpl) Is(locale, key string) bool {
	_, ok := s.translates[locale][key]
	return ok
}

// vKey formats a variable key for interpolation (prefix with colon)
func (s *ServiceImpl) vKey(key string) string {
	return ":" + key
}

// Translate retrieves a translation for the given locale and key.
// Falls back to the default locale if the translation is not found in the requested locale.
func (s *ServiceImpl) Translate(locale, key string) string {
	if v, ok := s.translates[locale][key]; ok {
		return v
	}
	if locale != s.locale {
		if v, ok := s.translates[s.locale][key]; ok {
			return v
		}
	}
	return key // Return the key itself if no translation is found
}

// T is a shorthand for Translate
func (s *ServiceImpl) T(locale, key string) string {
	return s.Translate(locale, key)
}

// Contains checks if a translation exists (with fallback to default locale)
func (s *ServiceImpl) Contains(locale, key string) bool {
	if _, ok := s.translates[locale][key]; ok {
		return true
	}
	if locale != s.locale {
		if _, ok := s.translates[s.locale][key]; ok {
			return true
		}
	}
	return false
}

// applyVariables replaces variable placeholders in a string with their values
func (s *ServiceImpl) applyVariables(value string, vars map[string]string) string {
	for k, v := range vars {
		value = strings.ReplaceAll(value, s.vKey(k), v)
	}
	return value
}

// Variables retrieves a translation and applies variable substitution
func (s *ServiceImpl) Variables(locale, key string, vars map[string]string) string {
	return s.applyVariables(s.Translate(locale, key), vars)
}

// V is a shorthand for Variables
func (s *ServiceImpl) V(locale, key string, vars map[string]string) string {
	return s.Variables(locale, key, vars)
}

// Choices retrieves a plural-aware translation based on a numeric value.
// Translation strings should be pipe-separated (e.g., "item|items" for English).
func (s *ServiceImpl) Choices(locale, key string, value int, vars map[string]string) string {
	result := s.Translate(locale, key)
	resultSplit := strings.Split(result, "|")
	resultSplitLen := len(resultSplit)

	if resultSplitLen < 2 {
		return result
	}

	// Handle negative values for pluralization
	if value < 0 {
		value = value * -1
	}

	// Determine the correct plural form based on locale-specific rules
	index := 0
	if locale == "ru" {
		index = s.choicesRuleRu(value, resultSplitLen)
	} else if locale == "en" {
		index = s.choicesRuleEn(value)
	}

	// Ensure we don't exceed the available choices
	resultSplitLen--
	if resultSplitLen >= index {
		result = resultSplit[index]
	}

	// Apply variable substitution if provided
	if vars != nil {
		return s.applyVariables(result, vars)
	}

	return result
}

// C is a shorthand for Choices
func (s *ServiceImpl) C(locale, key string, value int, vars map[string]string) string {
	return s.Choices(locale, key, value, vars)
}

// choicesRuleEn implements English pluralization rules:
// - 1: singular form (first choice)
// - other: plural form (second choice)
func (s *ServiceImpl) choicesRuleEn(value int) int {
	if value == 1 {
		return 0
	}
	return 1
}

// choicesRuleRu implements Russian pluralization rules with three forms:
// - 1, 21, 31... but not 11, 111...: singular
// - 2-4, 22-24, 32-34... but not 12-14: few
// - other: many
func (s *ServiceImpl) choicesRuleRu(value, choices int) int {
	if value%10 == 1 && value%100 != 11 {
		return 0
	}

	if choices == 2 || (value%10 >= 2 && value%10 <= 4 && (value%100 < 10 || value%100 >= 20)) {
		return 1
	}
	return 2
}

// TranslateValidationErrors translates validator.ValidationErrors to human-readable messages
// in the specified locale.
func (s *ServiceImpl) TranslateValidationErrors(locale string, e error) []errorsModule.FieldError {
	trans, _ := s.u.GetTranslator(locale)

	var validationErrors validator.ValidationErrors
	if errors.As(e, &validationErrors) {
		response := make([]errorsModule.FieldError, 0, len(validationErrors))
		for _, err := range validationErrors {
			response = append(response, errorsModule.FieldError{
				Name:    err.Field(),
				Message: err.Translate(trans),
			})
		}
		return response
	} else {
		// Log error if it's not a validation error
		errorsModule.E500(e, "translator.ServiceImpl.TranslateValidationErrors.As", "")
	}
	return nil
}

// TVE is a shorthand for TranslateValidationErrors
func (s *ServiceImpl) TVE(locale string, e error) []errorsModule.FieldError {
	return s.TranslateValidationErrors(locale, e)
}
