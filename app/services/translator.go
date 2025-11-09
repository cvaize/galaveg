package services

import (
	"encoding/json"
	"fmt"
	"galaveg/utils/flatten"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type TranslatorService struct {
	locale     string
	translates map[string]map[string]string
}

func NewTranslatorService(locale string, translates map[string]map[string]string) *TranslatorService {
	return &TranslatorService{locale, translates}
}

func NewTranslatorServiceFromFiles(dir, locale string) (*TranslatorService, error) {
	translates := map[string]map[string]string{}

	dir = filepath.Clean(dir)
	stat, err := os.Stat(dir)
	if err != nil || !stat.IsDir() {
		return nil, fmt.Errorf("NewTranslatorServiceFromFiles: translates folder not found: %s", dir)
	}

	var fullKey, lang, prefixKey, key, value string
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
				lang = ""
				prefixKey = ""

				if dotIndex == -1 {
					lang = fullKey
				} else {
					lang = fullKey[:dotIndex]
					prefixKey = fullKey[dotIndex+1:]
				}

				lang = strings.TrimSpace(lang)
				prefixKey = strings.TrimSpace(prefixKey)

				if lang == "" {
					return fmt.Errorf("the language is not defined by the path: %s", fullPath)
				}

				if _, ok := translates[lang]; !ok {
					translates[lang] = map[string]string{}
				}

				for key, valueAny = range flat {
					key = fmt.Sprintf("%s.%s", prefixKey, key)
					key = strings.TrimLeft(key, ".")
					key = strings.TrimSpace(key)
					value = strings.TrimSpace(fmt.Sprintf("%s", valueAny))
					translates[lang][key] = value
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("NewTranslatorServiceFromFiles: %w", err)
	}

	var variables, sp1, sp2 []string
	var varSt, v string
	var ok bool
	for lang, _ = range translates {
		for key, value = range translates[lang] {
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
					v, ok = translates[lang][strings.TrimSpace(varSt)]
					if !ok {
						v, ok = translates[locale][strings.TrimSpace(varSt)]
					}
					if ok {
						translates[lang][key] = strings.ReplaceAll(translates[lang][key], "{{"+varSt+"}}", v)
					}
				}
			}
		}
	}

	return NewTranslatorService(locale, translates), nil
}

func (s TranslatorService) Get(lang, key string) string {
	v, _ := s.translates[lang][key]
	return v
}

func (s TranslatorService) Is(lang, key string) bool {
	_, ok := s.translates[lang][key]
	return ok
}

func (s TranslatorService) vKey(key string) string {
	return ":" + key
}

func (s TranslatorService) Translate(lang, key string) string {
	if v, ok := s.translates[lang][key]; ok {
		return v
	}
	if lang != s.locale {
		if v, ok := s.translates[s.locale][key]; ok {
			return v
		}
	}
	return key
}

func (s TranslatorService) T(lang, key string) string {
	return s.Translate(lang, key)
}

func (s TranslatorService) Contains(lang, key string) bool {
	if _, ok := s.translates[lang][key]; ok {
		return true
	}
	if lang != s.locale {
		if _, ok := s.translates[s.locale][key]; ok {
			return true
		}
	}
	return false
}

func (s TranslatorService) applyVariables(value string, vars map[string]string) string {
	for k, v := range vars {
		value = strings.ReplaceAll(value, s.vKey(k), v)
	}
	return value
}

func (s TranslatorService) Variables(lang, key string, vars map[string]string) string {
	return s.applyVariables(s.Translate(lang, key), vars)
}

func (s TranslatorService) V(lang, key string, vars map[string]string) string {
	return s.Variables(lang, key, vars)
}

func (s TranslatorService) Choices(lang, key string, value int, vars map[string]string) string {
	result := s.Translate(lang, key)
	resultSplit := strings.Split(result, "|")
	resultSplitLen := len(resultSplit)

	if resultSplitLen < 2 {
		return result
	}

	if value < 0 {
		value = value * -1
	}

	index := 0
	if lang == "ru" {
		index = s.choicesRuleRu(value, resultSplitLen)
	} else if lang == "en" {
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

func (s TranslatorService) C(lang, key string, value int, vars map[string]string) string {
	return s.Choices(lang, key, value, vars)
}

func (s TranslatorService) choicesRuleEn(value int) int {
	if value > 1 {
		return 1
	} else {
		return 0
	}
}

func (s TranslatorService) choicesRuleRu(value, choices int) int {
	if value%10 == 1 && value%100 != 11 {
		return 0
	}

	if choices == 2 || (value%10 >= 2 && value%10 <= 4 && (value%100 < 10 || value%100 >= 20)) {
		return 1
	}
	return 2
}
