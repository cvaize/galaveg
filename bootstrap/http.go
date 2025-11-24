package bootstrap

import (
	"fmt"
	"galaveg/bootstrap/providers"
	"galaveg/config"
	"galaveg/routes"
	"galaveg/utils/logger"
	"galaveg/utils/path"
	"github.com/gin-gonic/gin"
	"strings"
)

func Http() *gin.Engine {
	cfg := config.MustDefault()
	ctx := providers.MustContext(cfg)
	ctx.TS.SetupValidator()

	defer ctx.Close()
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// TODO: Функции FuncMap вынести отдельно
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

	// TODO: Поместить "resources/html" конфигурацию
	templates := path.MustCollectFilepathBySuffix(cfg.GetFolder("resources/html"), ".gohtml")
	router.LoadHTMLFiles(templates...)

	if err := router.SetTrustedProxies(cfg.Http.AllowedHosts); err != nil {
		panic(err)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(middleware.CORSMiddleware())

	routes.Http(router, ctx)

	protocol := cfg.Http.Schema
	host := cfg.Http.Host
	port := cfg.Http.Port
	url := fmt.Sprintf("%s:%d", host, port)

	logger.Infof(fmt.Sprintf("Starting HTTP server at %s://%s", protocol, url))
	logger.Infof(fmt.Sprintf("Open HTTP website at %s", cfg.App.Url))
	if err := router.Run(url); err != nil {
		logger.Fatalf("Router Run() error: %s", err)
		panic(err)
	}

	return router
}
