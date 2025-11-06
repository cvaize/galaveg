package bootstrap

import (
	"fmt"
	"galaveg/bootstrap/singleton"
	"galaveg/routes"
	"galaveg/utils/logger"
	"github.com/gin-gonic/gin"
	"strings"
)

func Http() *gin.Engine {
	defer singleton.DB.Close()
	if singleton.C.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.FuncMap["eq"] = func(a, b string) bool { return a == b }
	router.FuncMap["ne"] = func(a, b string) bool { return a != b }
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
	// TODO: Проверить LoadHTMLGlob, возможно он не читает resources/html/*.gohtml . Если так, то создать отдельную функцию из filepath.WalkDir.
	router.LoadHTMLGlob("resources/html/**/*.gohtml")

	if err := router.SetTrustedProxies(singleton.C.App.AllowedHosts); err != nil {
		panic(err)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(middleware.CORSMiddleware())

	routes.Http(router)

	protocol := "http"
	host := singleton.C.App.Host
	port := singleton.C.App.Port
	url := fmt.Sprintf("%s:%d", host, port)
	logger.Infof(fmt.Sprintf("Starting HTTP server at %s://%s", protocol, url))
	if err := router.Run(url); err != nil {
		logger.Fatalf("Router Run() error: %s", err)
		panic(err)
	}

	return router
}
