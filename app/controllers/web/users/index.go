package users

import (
	"galaveg/app/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/auth", gin.H{
		"Lang":    "ru",
		"Title":   "Вход - Galaveg",
		"Heading": "Вход",
		"DarkMode": map[string]any{
			"DarkTitle":  "1",
			"LightTitle": "2",
			"AutoTitle":  "3",
			"Value":      "auto",
		},
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
		"Locale": map[string]string{
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
