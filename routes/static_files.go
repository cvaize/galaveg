package routes

import (
	"galaveg/app/controllers/api/v1"
	"github.com/gin-gonic/gin"
)

func staticFilesRegister(router *gin.Engine) {
	var ctrl v1.Controller
	css := router.Group("/css")
	css.GET("/app.css", ctrl.Index)

	js := router.Group("/js")
	js.GET("/app.js", ctrl.Index)

	svg := router.Group("/svg")
	svg.GET("/logo.svg", ctrl.Index)
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
