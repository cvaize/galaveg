package dto

type Alert struct {
	Style   string
	Content string
}

func NewInfoAlert(content string) Alert {
	return Alert{Style: "info", Content: content}
}

func NewSuccessAlert(content string) Alert {
	return Alert{Style: "success", Content: content}
}

func NewWarningAlert(content string) Alert {
	return Alert{Style: "warning", Content: content}
}

func NewErrorAlert(content string) Alert {
	return Alert{Style: "error", Content: content}
}
