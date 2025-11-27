package contracts

import (
	"galaveg/app/dto"
	"galaveg/app/dto/mailables"
	"galaveg/bootstrap/providers"
)

// TODO: imports galaveg/bootstrap/providers.Context VS galaveg/app/contracts
type NotificationContract interface {
	IsQueue(ctx *providers.Context) bool
	// Via Get the notification's delivery channels. Example: []string{"mail"}
	Via(ctx *providers.Context) []string
	BuildEmailMessage(ctx *providers.Context) (*mailables.EmailMessage, *dto.Error)
}
