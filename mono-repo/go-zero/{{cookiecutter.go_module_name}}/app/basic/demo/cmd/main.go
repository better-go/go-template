package main

import (
	"github.com/better-go/pkg/x/go-zero/cmd"

	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/server"
)

func main() {
	// auto select:
	cmd.Runner(server.NewServerSelector())
}
