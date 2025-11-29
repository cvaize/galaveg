package routes

import (
	"galaveg/internal/bootstrap/http/handlers/http/controllers/static_files"
	"github.com/gin-gonic/gin"
)

func staticFilesRouter(r *gin.Engine) {
	g := r.Group("/static")
	//g.Use(middlewares.GzipStaticMiddleware())
	g.Static("/", "./public")

	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.GET("/css/app.css", static_files.Gzip("./public/css/app.min.css", "text/css; charset=utf-8"))
	r.GET("/js/app.js", static_files.Gzip("./public/js/app.min.js", "application/javascript; charset=utf-8"))
	r.GET("/svg/logo.svg", static_files.Gzip("./public/svg/logo.svg", "image/svg+xml"))

	//	cfg.service(
	//		web::resource("/storage/private-files/{filename}")
	//	.wrap(WebAuthMiddleware)
	//	.route(web::get().to(static_files::storage::private)),
	//);
}
