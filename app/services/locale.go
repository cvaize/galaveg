package services

import (
	"errors"
	"galaveg/app/dto"
)

type LocaleService struct {
	localeCookieKey string
	defaultLocale   dto.Locale
	localesMap      map[string]dto.Locale
	localesCodes    []string
	locales         []dto.Locale
}

func NewLocaleService(defaultLocaleCode, localeCookieKey string, locales []dto.Locale) (*LocaleService, error) {
	localesMap := map[string]dto.Locale{}
	var localesCodes []string

	for _, l := range locales {
		localesCodes = append(localesCodes, l.Code)
		localesMap[l.Code] = l
	}
	defaultLocale, ok := localesMap[defaultLocaleCode]

	if !ok {
		return nil, errors.New("default locale not found")
	}

	return &LocaleService{
		localeCookieKey,
		defaultLocale,
		localesMap,
		localesCodes,
		locales,
	}, nil
}

func MustLocaleService(defaultLocaleCode, localeCookieKey string, locales []dto.Locale) *LocaleService {
	s, e := NewLocaleService(defaultLocaleCode, localeCookieKey, locales)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *LocaleService) GetLocale(code string) dto.Locale {
	l, _ := s.localesMap[code]
	return l
}

func (s *LocaleService) GetLocaleOrDefault(code string) dto.Locale {
	l, ok := s.localesMap[code]
	if ok {
		return l
	}
	return s.defaultLocale
}

func (s *LocaleService) GetDefault() dto.Locale {
	return s.defaultLocale
}

func (s *LocaleService) GetLocales() []dto.Locale {
	return s.locales
}

func (s *LocaleService) existsCode(code string) string {
	_, ok := s.localesMap[code]
	if ok {
		return code
	}
	return s.defaultLocale.Code
}

func (s *LocaleService) GetLocalesCodes() []string {
	return s.localesCodes
}

//pub fn get_locale_code(&self, req: Option<&HttpRequest>, user: Option<&User>) -> String {
//if let Some(req) = req {
//if let Some(locale) = req.cookie(&self.config.app.locale_cookie_key) {
//let locale = locale.value().to_string();
//if StrMinMaxLength::apply(&locale, 1, 6) {
//return self.exists_locale_code_or_default(locale);
//}
//}
//}
//if let Some(user) = user {
//if let Some(locale) = &user.locale {
//if StrMinMaxLength::apply(&locale, 1, 6) {
//return self.exists_locale_code_or_default(locale.to_string());
//}
//}
//}
//if let Some(req) = req {
//if let Some(header) = req.headers().get(ACCEPT_LANGUAGE) {
//let languages = accept_language::intersection(
//header.to_str().unwrap_or(&self.config.app.locale),
//&self.get_locales_codes_ref(),
//);
//
//if let Some(locale) = languages.first() {
//return self.exists_locale_code_or_default(locale.to_string());
//}
//}
//}
//
//self.config.app.locale.to_string()
//}
