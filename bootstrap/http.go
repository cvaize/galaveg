package bootstrap

import (
	"fmt"
	"galaveg/config"
	"galaveg/connections"
	"galaveg/routes"
	"galaveg/utils/logger"
	"github.com/gin-gonic/gin"
)

func Http() *gin.Engine {
	defer connections.DB.Close()
	if config.Config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.FuncMap["eq"] = func(a, b string) bool { return a == b }
	router.FuncMap["ne"] = func(a, b string) bool { return a != b }
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
	router.LoadHTMLGlob("resources/html/**/*.gohtml")

	if err := router.SetTrustedProxies(config.Config.App.AllowedHosts); err != nil {
		panic(err)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(middleware.CORSMiddleware())

	routes.Http(router)

	protocol := "http"
	host := config.Config.App.Host
	port := config.Config.App.Port
	url := fmt.Sprintf("%s:%d", host, port)
	logger.Infof(fmt.Sprintf("Starting HTTP server at %s://%s", protocol, url))
	if err := router.Run(url); err != nil {
		logger.Fatalf("Router Run() error: %s", err)
		panic(err)
	}

	return router
}
