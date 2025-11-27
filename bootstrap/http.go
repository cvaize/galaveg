package bootstrap

import (
	"fmt"
	"galaveg/bootstrap/providers"
	"galaveg/config"
	"galaveg/routes"
	"galaveg/utils/logger"
	"github.com/gin-gonic/gin"
)

func Http() *gin.Engine {
	cfg := config.MustDefault()
	ctx := providers.MustContext(cfg)

	defer ctx.Close()
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.SetHTMLTemplate(ctx.Infra.Html)

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
