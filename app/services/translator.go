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

	var fullKey string
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
				dotIndex := strings.Index(fullKey, ".")
				lang := ""
				prefixKey := ""

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

				for key, valueAny := range flat {
					key = fmt.Sprintf("%s.%s", prefixKey, key)
					key = strings.TrimLeft(key, ".")
					key = strings.TrimSpace(key)
					value := strings.TrimSpace(fmt.Sprintf("%s", valueAny))
					translates[lang][key] = value
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("NewTranslatorServiceFromFiles: %w", err)
	}

	fmt.Println(translates)

	return NewTranslatorService(locale, translates), nil
}
