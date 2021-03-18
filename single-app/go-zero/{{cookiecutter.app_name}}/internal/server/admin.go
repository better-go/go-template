package server

import (
	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/router"
	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/service/admin"
	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

// internal admin HTTP/API server:
type AdminServer struct{}

func (m *AdminServer) Run(configFile string) {
	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)

	// admin api:
	svc := admin.NewService(cfg)
	defer svc.Close()

	// new engine:
	engine := rest.MustNewServer(cfg.Server.Admin.RestConf)
	defer engine.Stop()

	// register routes:
	router.RegisterAdminRoutes(engine, svc)

	log.Infof("starting admin api engine at %v\n", cfg.Server.Admin.RestConf)
	engine.Start()
}
