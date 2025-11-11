package users

import (
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	//page.users.index.header
	c.HTML(http.StatusOK, "layouts/list", gin.H{
		"Locale":   "ru",
		"Title":    "Главная - Galaveg",
		"Heading":  singleton.TS.T("en", "page.users.index.header"),
		"DarkMode": "auto",
		"Breadcrumbs": []map[string]string{
			{
				"Label": "Главная",
				"Href":  "/",
			},
			{
				"Label": "Пользователи",
				"Href":  "/users",
			},
		},
	})
}
