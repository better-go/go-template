package http

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"

    "github.com/better-go/pkg/log"
)

// 极光推送 API 对接:
type DemoHttp struct {
}

func newDemoHttp(cfg *config.HttpUnit) *DemoHttp {

    return &DemoHttp{
    }
}

func (m *DemoHttp) Hello(message string) error {
    log.Debugf("demo.Hello done, req=%+v", message)
    return nil
}
