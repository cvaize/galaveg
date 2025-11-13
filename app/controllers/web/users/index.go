package users

import (
	"fmt"
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	fmt.Println(c.FullPath())
	//page.users.index.header
	c.HTML(http.StatusOK, "layouts/list", gin.H{
		"Lang":     "ru",
		"Title":    "Главная - Galaveg",
		"Heading":  singleton.TS.T("en", "page.users.index.header"),
		"DarkMode": "auto",
		"Breadcrumbs": []map[string]string{
			{
				"Text": "Главная",
				"Href": "/",
			},
			{
				"Text": "Пользователи",
				"Href": "/users",
			},
		},
	})
}
