package cache

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

/*
所有 HTTP 接口对接

*/

type Dao struct {
    // biz group:
    Demo *DemoCache // demo cache
}

func New(cfg *config.CacheUnit) *Dao {
    return &Dao{
        Demo: newDemoCache(cfg),
    }
}

func (m *Dao) Close() {
    return
}
