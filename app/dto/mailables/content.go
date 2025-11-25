package mailables

type Content struct {
	Text       string
	HtmlString string
	Markdown   string
	// TODO: Implement rendering in deferred message sending
	//View       string
	//With       any
}
