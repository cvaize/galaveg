package serve

import (
	"galaveg/internal/bootstrap/chat"
	"galaveg/internal/bootstrap/http"
)

func Run() {
	chat.Run()
	http.Run()
}
