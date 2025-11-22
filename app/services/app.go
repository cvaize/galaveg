package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

type AppService struct {
	cfg *config.Config
	ls  *LocaleService
	rs  *RoleService
	ts  *TranslatorService
	url *url.URL
}

func NewAppService(c *config.Config, ls *LocaleService, rs *RoleService, ts *TranslatorService) (*AppService, error) {
	u, err := url.Parse(c.App.Url)
	if err != nil {
		panic(err)
	}
	return &AppService{c, ls, rs, ts, u}, nil
}

func MustAppService(c *config.Config, ls *LocaleService, rs *RoleService, ts *TranslatorService) *AppService {
	s, e := NewAppService(c, ls, rs, ts)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *AppService) DarkMode(c *gin.Context) string {
	var val string
	if c != nil {
		val, _ = c.Cookie(s.cfg.App.DarkModeCookieKey)
	}
	if val != "auto" && val != "dark" && val != "light" {
		val = "auto"
	}

	return val
}

func (s *AppService) Locale(c *gin.Context, user *dto.User) string {
	var val string
	if c != nil {
		// manually selected by the user in the browser
		val, _ = c.Cookie(s.cfg.App.LocaleCookieKey)
		if val != "" {
			if s.ls.Exists(val) {
				return val
			} else {
				return s.cfg.App.Locale
			}
		}
	}
	if user != nil && user.Locale != "" {
		// the locale is selected from the database
		if s.ls.Exists(user.Locale) {
			return user.Locale
		} else {
			return s.cfg.App.Locale
		}
	}
	if c != nil {
		// Accept-Language: *
		// Accept-Language: fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5
		val = c.GetHeader("Accept-Language")
		if val != "" {
			if index := strings.IndexAny(val, "-,;"); index != -1 {
				if s.ls.Exists(val[:index]) {
					return val[:index]
				} else {
					return s.cfg.App.Locale
				}
			}
		}
	}
	return s.cfg.App.Locale
}

func (s *AppService) Csrf(c *gin.Context) string {
	return ""
}

func (s *AppService) Alerts(c *gin.Context) []dto.Alert {
	return []dto.Alert{}
}

func (s *AppService) Url() string {
	return s.url.String()
}

func (s *AppService) CloneUrl() url.URL {
	return *s.url
}

// TODO: Вынести WebDataCtx, NewWebDataCtx, GetWebData в app/view/layouts или app/view/components

type WebDataCtx struct {
	User     *dto.User
	Locale   dto.Locale
	Locales  []dto.Locale
	Alerts   []dto.Alert
	SiteUrl  string
	Path     string
	Heading  string
	DarkMode string
	Csrf     string
}

func (s *AppService) NewWebDataCtx(c *gin.Context) (WebDataCtx, error) {
	locale := s.ls.GetLocale(s.Locale(c, nil))
	locales := s.ls.GetLocales()

	return WebDataCtx{
		User:     nil,
		Locale:   locale,
		Locales:  locales,
		Alerts:   s.Alerts(c),
		SiteUrl:  s.Url(),
		Path:     c.FullPath(),
		DarkMode: s.DarkMode(c),
		Csrf:     s.Csrf(c),
	}, nil
}

func (s *AppService) GetWebData(ctx *WebDataCtx) gin.H {
	//let mut sidebar_users_index: Option<String> = None;
	//let mut sidebar_roles_index: Option<String> = None;
	//let mut sidebar_files: Option<String> = None;
	//let mut is_sidebar_users_dropdown = false;
	//
	//if let Ok(roles) = role_service.all() {
	//	let is_users_show = UserPolicy::can_show(user, &roles);
	//	if is_users_show {
	//		sidebar_users_index =
	//			Some(translator_service.translate(lang, "layout.sidebar.users.index"));
	//	}
	//	let is_roles_show = RolePolicy::can_show(user, &roles);
	//	if is_roles_show {
	//		sidebar_roles_index =
	//			Some(translator_service.translate(lang, "layout.sidebar.users.roles"));
	//	}
	//	is_sidebar_users_dropdown = is_users_show && is_roles_show;
	//
	//	if FilePolicy::can_show(user, &roles) {
	//		sidebar_files = Some(translator_service.translate(lang, "layout.sidebar.files"));
	//	}
	//}

	//	"sidebar": {
	//	"home": translator_service.translate(lang, "layout.sidebar.home"),
	//		"users": {
	//		"is_dropdown": is_sidebar_users_dropdown,
	//			"index": sidebar_users_index,
	//			"roles": sidebar_roles_index,
	//	},
	//	"files": sidebar_files,
	//		"profile": translator_service.translate(lang, "layout.sidebar.profile"),
	//		"logout": translator_service.translate(lang, "layout.sidebar.logout"),
	//},
	return gin.H{
		"Lang":     ctx.Locale.Code,
		"Brand":    s.ts.T(ctx.Locale.Code, "layout.brand"),
		"Title":    s.ts.T(ctx.Locale.Code, "app.name"),
		"DarkMode": ctx.DarkMode,
		"Csrf":     ctx.Csrf,
		// "http://localhost:3000/"
		"SiteUrl": ctx.SiteUrl,
		// "/" or "/users/"
		"Path":    ctx.Path,
		"Alerts":  ctx.Alerts,
		"User":    ctx.User,
		"Locale":  ctx.Locale,
		"Locales": ctx.Locales,
	}
}
