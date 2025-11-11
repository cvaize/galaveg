package auth

import (
	"galaveg/app/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "layouts/auth", gin.H{
		"Locale":   "ru",
		"Title":    "Вход - Galaveg",
		"Heading":  "Вход",
		"DarkMode": "auto",
		"Form": map[string]any{
			"Action": "/login",
			"Method": "post",
			"Fields": []map[string]any{
				{
					"Label":  "E-mail",
					"Type":   "email",
					"Name":   "email",
					"Value":  "",
					"Errors": []string{},
				},
				{
					"Label":  "Пароль",
					"Type":   "password",
					"Name":   "password",
					"Value":  "",
					"Errors": []string{},
				},
			},
			"Submit": map[string]string{
				"Label": "Войти",
			},
			"ResetPassword": map[string]string{
				"Label": "Сбросить пароль?",
				"Href":  "/reset-password",
			},
			"Register": map[string]string{
				"Label": "Зарегистрироваться",
				"Href":  "/register",
			},
			"Errors": []string{},
		},
		"locale": map[string]string{
			"FullName":  "ru",
			"ShortName": "ru",
			"Code":      "ru",
		},
		"Locales": []map[string]string{
			{
				"FullName":  "en",
				"ShortName": "en",
				"Code":      "en",
			},
		},
		"Back": map[string]string{
			"Href":  "/login",
			"Label": "Назад",
		},
		"Alerts": []dto.Alert{
			{"success", "Test1"},
			{"warning", "Test2"},
		},
	})
}
