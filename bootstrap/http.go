package bootstrap

import (
	"database/sql"
	"fmt"
	"galaveg/app/dto"
	"galaveg/app/services"
	"galaveg/bootstrap/providers"
	"galaveg/config"
	"galaveg/connections/db"
	"galaveg/routes"
	"galaveg/utils/logger"
	"galaveg/utils/path"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path/filepath"

	"strings"
)

func httpProvideContext(C *config.Config, DB *sql.DB) *providers.Context {
	TS := services.MustTranslatorServiceFromFiles(C.GetFolder("resources/translates/"), C.App.Locale)
	LS := services.MustLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	})
	RS := services.MustRoleService()
	AS := services.MustAppService(C.App, LS, RS, TS)
	return &providers.Context{
		C:  C,
		TS: TS,
		LS: LS,
		RS: RS,
		AS: AS,
	}
}

func Http() *gin.Engine {
	C := config.New(filepath.Join(viper.GetString("APP_FOLDER"), ".env"))
	DB := db.New(C.Db)
	ctx := httpProvideContext(C, DB)

	defer DB.Close()
	if C.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.FuncMap["eq"] = func(a, b string) bool { return a == b }
	router.FuncMap["eqInt"] = func(a, b int) bool { return a == b }
	router.FuncMap["ne"] = func(a, b string) bool { return a != b }
	router.FuncMap["neInt"] = func(a, b int) bool { return a != b }
	router.FuncMap["replace"] = strings.ReplaceAll
	router.FuncMap["unless"] = func(v any) bool {
		return v == nil || v == "" || v == false
	}
	router.FuncMap["dict"] = func(values ...interface{}) map[string]interface{} {
		dict := make(map[string]interface{})
		for i := 0; i < len(values); i += 2 {
			key := values[i].(string)
			value := values[i+1]
			dict[key] = value
		}
		return dict
	}
	router.FuncMap["sub1"] = func(x int) int { return x - 1 }
	router.FuncMap["startsWith"] = func(s, prefix string) bool {
		return strings.HasPrefix(s, prefix)
	}

	templates := path.MustCollectFilepathBySuffix(C.GetFolder("resources/html"), ".gohtml")
	router.LoadHTMLFiles(templates...)

	if err := router.SetTrustedProxies(C.App.AllowedHosts); err != nil {
		panic(err)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(middleware.CORSMiddleware())

	routes.Http(router, ctx)

	protocol := "http"
	host := C.App.Host
	port := C.App.Port
	url := fmt.Sprintf("%s:%d", host, port)
	logger.Infof(fmt.Sprintf("Starting HTTP server at %s://%s", protocol, url))
	if err := router.Run(url); err != nil {
		logger.Fatalf("Router Run() error: %s", err)
		panic(err)
	}

	return router
}
