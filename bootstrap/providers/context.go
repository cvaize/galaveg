package providers

import (
	"galaveg/app/services"
	"galaveg/config"
)

type Context struct {
	C  *config.Config
	TS *services.TranslatorService
	LS *services.LocaleService
	AS *services.AppService
	RS *services.RoleService
}
