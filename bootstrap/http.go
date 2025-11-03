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

	router.LoadHTMLGlob("resources/html/**/*")

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
