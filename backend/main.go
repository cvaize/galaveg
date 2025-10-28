package main

import (
	"galaveg/bootstrap"
	"os"
)

func main() {
	bootstrap.Root()

	arg := ""

	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "server":
		bootstrap.Server()
		break
	default:
		bootstrap.Console()
		break
	}
}
