package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "layouts/home", gin.H{
		"Lang":     "ru",
		"Title":    "Главная - Galaveg",
		"Heading":  "Компоненты",
		"DarkMode": "auto",
		"Breadcrumbs": []map[string]string{
			{
				"Label": "Главная",
				"Href":  "/",
			},
		},
	})
}
