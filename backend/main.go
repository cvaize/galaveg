/*
Copyright © 2025 Dmitry Orlov cvaize@gmail.com
*/
package main

import (
	"galaveg/bootstrap"
	"galaveg/cmd"
)

func main() {
	bootstrap.Config()
	cmd.Execute()
}
