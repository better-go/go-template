package http

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

/*
所有 HTTP 接口对接

*/

type Dao struct {
    // biz group:
    Demo *DemoHttp // 极光推送服务
}

func New(cfg *config.HTTPConfig) *Dao {
    return &Dao{
        Demo: newDemoHttp(cfg),
    }
}

func (m *Dao) Close() {
    return
}
