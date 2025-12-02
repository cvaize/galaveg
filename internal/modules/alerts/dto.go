package alerts

import "encoding/gob"

type AlertDto struct {
	Style   string
	Content string
}

func init() {
	gob.Register(AlertDto{})
	gob.Register([]AlertDto{})
}

func NewInfoAlert(content string) AlertDto {
	return AlertDto{Style: "info", Content: content}
}

func NewSuccessAlert(content string) AlertDto {
	return AlertDto{Style: "success", Content: content}
}

func NewWarningAlert(content string) AlertDto {
	return AlertDto{Style: "warning", Content: content}
}

func NewErrorAlert(content string) AlertDto {
	return AlertDto{Style: "error", Content: content}
}
