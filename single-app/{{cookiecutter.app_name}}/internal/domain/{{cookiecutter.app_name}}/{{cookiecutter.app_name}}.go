package {{cookiecutter.app_name}}

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

// 推送服务业务逻辑
type Scope struct {
    // inner global use:
    g *dao.MetaResource

    // biz:
    Demo *DemoScope
}

//
func NewScope(httpCfg *config.HTTPConfig) *Scope {
    // global:
    g := dao.NewMetaResource(httpCfg)

    return &Scope{
        g: g,

        // biz:
        Demo: newDemoScope(g),
    }
}

func (m *Scope) Close() {
    m.g.Close()
    return
}
