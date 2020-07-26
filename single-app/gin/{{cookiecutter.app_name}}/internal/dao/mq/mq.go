package mq

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

/*
所有 HTTP 接口对接

*/

type Dao struct {
    // biz group:
    Demo *DemoMQ // demo http
}

func New(cfg *config.MQUnit) *Dao {
    return &Dao{
        Demo: newDemoMQ(cfg),
    }
}

func (m *Dao) Close() {
    return
}
