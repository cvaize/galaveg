package app

import (
	"galaveg/config"
	"galaveg/routes"
	"galaveg/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Create() *gin.Engine {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
		panic(err)
	}

	viper.SetDefault("APP_LOG_LEVEL", "info")
	logger.SetLogLevel(viper.GetString("APP_LOG_LEVEL"))

	environment := viper.GetBool("APP_DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("APP_ALLOWED_HOSTS")
	router := gin.New()

	if err := router.SetTrustedProxies([]string{allowedHosts}); err != nil {
		logger.Fatalf("router SetTrustedProxies() error: %s", err)
		panic(err)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(middleware.CORSMiddleware())

	routes.Register(router)

	logger.Infof("Starting HTTP server at http://0.0.0.0:8080")
	if err := router.Run("127.0.0.1:8080"); err != nil {
		logger.Fatalf("router Run() error: %s", err)
	}

	return router
}
