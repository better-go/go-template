package main

import (
	"github.com/better-go/pkg/x/go-zero/cmd"

	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/server"
)

func main() {
	// auto select:
	cmd.Runner(server.NewServerSelector())
}
