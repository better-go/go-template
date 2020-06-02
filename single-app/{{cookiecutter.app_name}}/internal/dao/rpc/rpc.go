package rpc

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

/*
所有 HTTP 接口对接

*/

type Dao struct {
    // biz group:
    Demo *DemoRpc // demo rpc
}

func New(cfg *config.RpcUnit) *Dao {
    return &Dao{
        Demo: newDemoRpc(cfg),
    }
}

func (m *Dao) Close() {
    return
}
