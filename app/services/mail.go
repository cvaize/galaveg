package services

import "galaveg/config"

// https://github.com/wneessen/go-mail/wiki/Getting-started#installation

type MailService struct {
	config config.AppConfig
}
