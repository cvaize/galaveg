package routes

import (
	"galaveg/app/controllers/static_files"
	//"galaveg/app/middlewares"
	"github.com/gin-gonic/gin"
)

func staticFilesRegister(r *gin.Engine) {
	g := r.Group("/static")
	//g.Use(middlewares.GzipStaticMiddleware())
	g.Static("/", "./public")

	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.GET("/app.css", static_files.Gzip("./public/css/app.min.css", "text/css; charset=utf-8"))
	r.GET("/app.js", static_files.Gzip("./public/js/app.min.js", "application/javascript; charset=utf-8"))
	r.GET("/logo.svg", static_files.Gzip("./public/svg/logo.svg", "image/svg+xml"))

	//	cfg.service(
	//		web::resource("/storage/private-files/{filename}")
	//	.wrap(WebAuthMiddleware)
	//	.route(web::get().to(static_files::storage::private)),
	//);
}
