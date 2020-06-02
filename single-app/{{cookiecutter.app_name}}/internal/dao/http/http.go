package http

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

/*
所有 HTTP 接口对接

*/

type Dao struct {
    // biz group:
    Demo *DemoHttp // demo http
}

func New(cfg *config.HttpUnit) *Dao {
    return &Dao{
        Demo: newDemoHttp(cfg),
    }
}

func (m *Dao) Close() {
    return
}
