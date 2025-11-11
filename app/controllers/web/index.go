package web

import (
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	darkMode := singleton.AS.DarkMode(c)
	locale := singleton.AS.Locale(c, nil)

	c.HTML(http.StatusOK, "layouts/home", gin.H{
		"Locale":   locale,
		"Title":    "Главная - Galaveg",
		"Heading":  "Компоненты",
		"DarkMode": darkMode,
		"Breadcrumbs": []map[string]string{
			{
				"Label": "Главная",
				"Href":  "/",
			},
		},
	})
}
