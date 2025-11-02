package routes

import (
	"galaveg/app/controllers/static_files"
	//"galaveg/app/middlewares"
	"github.com/gin-gonic/gin"
)

func staticFilesRegister(r *gin.Engine) {
	g := r.Group("/static")
	//g.Use(middlewares.GzipStaticMiddleware())
	g.Static("/", "./static")

	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.GET("/app.css", static_files.Gzip("./static/css/app.min.css", "text/css; charset=utf-8"))
	r.GET("/app.js", static_files.Gzip("./static/js/app.min.js", "application/javascript; charset=utf-8"))
	r.GET("/logo.svg", static_files.Gzip("./static/svg/logo.svg", "image/svg+xml"))

	//	cfg.service(
	//		web::resource("/storage/files/{filename}")
	//	.route(web::get().to(static_files::storage::public)),
	//);
	//	cfg.service(
	//		web::resource("/storage/private-files/{filename}")
	//	.wrap(WebAuthMiddleware)
	//	.route(web::get().to(static_files::storage::private)),
	//);
}
