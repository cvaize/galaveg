package routes

import (
	"galaveg/app/middlewares"
	"github.com/gin-gonic/gin"
)

func staticFilesRegister(r *gin.Engine) {
	g := r.Group("/static")
	g.Use(middlewares.GzipStaticMiddleware())
	g.Static("/", "./static")

	r.StaticFile("/favicon.ico", "./static/favicon.ico")

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
