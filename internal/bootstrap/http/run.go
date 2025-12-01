package http

import (
	"fmt"
	"galaveg/internal/bootstrap/http/context"
	"galaveg/internal/bootstrap/http/handlers/http/routes"
	"galaveg/internal/config"
	"galaveg/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.MustDefault()
	logger.SetLogLevel(cfg.App.LogLevel)
	ctx := context.Must(cfg)

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

	routes.Router(router, ctx)

	protocol := cfg.Http.Schema
	host := cfg.Http.Host
	port := cfg.Http.Port
	url := fmt.Sprintf("%s:%d", host, port)

	logger.Infof(fmt.Sprintf("Starting HTTP server at %s://%s", protocol, url))
	logger.Infof(fmt.Sprintf("Open HTTP website at %s", cfg.Http.Url))
	if err := router.Run(url); err != nil {
		logger.Fatalf("Router Run() error: %s", err)
		panic(err)
	}
}
