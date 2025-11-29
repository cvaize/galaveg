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

type Service struct {
	locale     string
	translates map[string]map[string]string
	u          *ut.UniversalTranslator
}

func NewService(locale string, translates map[string]map[string]string) *Service {
	enLocale := en.New()
	ruLocale := ru.New()

	u := ut.New(enLocale, enLocale, ruLocale)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		transEN, _ := u.GetTranslator("en")
		_ = enTranslations.RegisterDefaultTranslations(v, transEN)

		transRU, _ := u.GetTranslator("ru")
		_ = ruTranslations.RegisterDefaultTranslations(v, transRU)
	}
	return &Service{locale, translates, u}
}

func NewServiceFromFiles(dir, locale string) (*Service, *errorsModule.Error) {
	translates := map[string]map[string]string{}

	dir = filepath.Clean(dir)
	stat, err := os.Stat(dir)
	if err != nil || !stat.IsDir() {
		return nil, errorsModule.E500(err, "translator.Service.NewServiceFromFiles", fmt.Sprintf("translates folder not found: %s", dir))
	}

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
		if strings.HasSuffix(strings.ToLower(d.Name()), ".json") {
			//
			fullKey = strings.ReplaceAll(fullPath, dir, "")
			fullKey = strings.ReplaceAll(fullKey, string(filepath.Separator), ".")
			fullKey = strings.TrimLeft(fullKey, ".")
			fullKey = strings.ReplaceAll(fullKey, ".json", "")
			content, err := os.ReadFile(fullPath)
			if err != nil {
				return err
			}

			var parsed map[string]interface{}
			if err := json.Unmarshal(content, &parsed); err != nil {
				return err
			}

			flat, err := flatten.Flatten(parsed, "", flatten.DotStyle)
			if err != nil {
				return err
			}

			if len(flat) != 0 {
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

				if _, ok := translates[locale1]; !ok {
					translates[locale1] = map[string]string{}
				}

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
		return nil, errorsModule.E500(err, "translator.Service.NewServiceFromFiles", "")
	}

	var variables, sp1, sp2 []string
	var varSt, v string
	var ok bool
	for locale1, _ = range translates {
		for key, value = range translates[locale1] {
			variables = []string{}

			sp1 = strings.Split(value, "{{")
			for _, varSt = range sp1 {
				sp2 = strings.Split(varSt, "}}")

				if len(sp2) > 1 {
					variables = append(variables, sp2[0])
				}
			}

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

	return NewService(locale, translates), nil
}

func (s *Service) GetLocale() string {
	return s.locale
}

func (s *Service) GetTranslates() map[string]map[string]string {
	return s.translates
}

func (s *Service) Get(locale, key string) string {
	v, _ := s.translates[locale][key]
	return v
}

func (s *Service) Is(locale, key string) bool {
	_, ok := s.translates[locale][key]
	return ok
}

func (s *Service) vKey(key string) string {
	return ":" + key
}

func (s *Service) Translate(locale, key string) string {
	if v, ok := s.translates[locale][key]; ok {
		return v
	}
	if locale != s.locale {
		if v, ok := s.translates[s.locale][key]; ok {
			return v
		}
	}
	return key
}

func (s *Service) T(locale, key string) string {
	return s.Translate(locale, key)
}

func (s *Service) Contains(locale, key string) bool {
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

func (s *Service) applyVariables(value string, vars map[string]string) string {
	for k, v := range vars {
		value = strings.ReplaceAll(value, s.vKey(k), v)
	}
	return value
}

func (s *Service) Variables(locale, key string, vars map[string]string) string {
	return s.applyVariables(s.Translate(locale, key), vars)
}

func (s *Service) V(locale, key string, vars map[string]string) string {
	return s.Variables(locale, key, vars)
}

func (s *Service) Choices(locale, key string, value int, vars map[string]string) string {
	result := s.Translate(locale, key)
	resultSplit := strings.Split(result, "|")
	resultSplitLen := len(resultSplit)

	if resultSplitLen < 2 {
		return result
	}

	if value < 0 {
		value = value * -1
	}

	index := 0
	if locale == "ru" {
		index = s.choicesRuleRu(value, resultSplitLen)
	} else if locale == "en" {
		index = s.choicesRuleEn(value)
	}

	resultSplitLen--
	if resultSplitLen >= index {
		result = resultSplit[index]
	}

	if vars != nil {
		return s.applyVariables(result, vars)
	}

	return result
}

func (s *Service) C(locale, key string, value int, vars map[string]string) string {
	return s.Choices(locale, key, value, vars)
}

func (s *Service) choicesRuleEn(value int) int {
	if value == 1 {
		return 0
	}
	return 1
}

func (s *Service) choicesRuleRu(value, choices int) int {
	if value%10 == 1 && value%100 != 11 {
		return 0
	}

	if choices == 2 || (value%10 >= 2 && value%10 <= 4 && (value%100 < 10 || value%100 >= 20)) {
		return 1
	}
	return 2
}

func (s *Service) TranslateValidationErrors(locale string, e error) []errorsModule.FieldError {
	trans, _ := s.u.GetTranslator(locale)

	var validationErrors validator.ValidationErrors
	if errors.As(e, &validationErrors) {
		response := make([]errorsModule.FieldError, 0, len(validationErrors))
		for _, err := range validationErrors {
			response = append(response, errorsModule.FieldError{Name: err.Field(), Message: err.Translate(trans)})
		}
		return response
	} else {
		//goland:noinspection GoUnhandledErrorResult
		errorsModule.E500(e, "translator.Service.TranslateValidationErrors.As", "")
	}
	return nil
}

func (s *Service) TVE(locale string, e error) []errorsModule.FieldError {
	return s.TranslateValidationErrors(locale, e)
}
