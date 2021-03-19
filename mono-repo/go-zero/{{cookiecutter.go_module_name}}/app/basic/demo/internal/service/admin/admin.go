package admin

import (
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/domain/demo"
	"{{cookiecutter.go_module_name}}/app/basic/demo/proto/config"
)

/*
	Admin API Service: 提供内部 API

*/
type Service struct {
	d *queue.Domain // 引入业务单元
}

func NewService(cfg config.Config) *Service {
	return &Service{
		d: queue.NewDomain(cfg, false),
	}
}

func (m *Service) Close() {
	m.d.Close()
}
